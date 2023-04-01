package classify

import (
	"401-company/internal/service"
	"401-company/model"
	"401-company/model/dao"
	"401-company/model/do"
	"401-company/model/entity"
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sClassify struct {
	conf gdb.CacheOption
}

func init() {
	service.RegisterClassify(New())
}

func New() *sClassify {
	return &sClassify{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

func (s *sClassify) CreateClassify(ctx context.Context, info model.Classify) (*entity.Classify, error) {
	info.Id = 0
	return s.SaveClassify(ctx, info)
}

func (s *sClassify) UpdateClassify(ctx context.Context, info model.Classify) (*entity.Classify, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新分类信息ID参数错误"), "", dao.Classify.Table())
	}
	return s.SaveClassify(ctx, info)
}

func (s *sClassify) SaveClassify(ctx context.Context, info model.Classify) (*entity.Classify, error) {
	classifyInfo := entity.Classify{
		Id:       info.Id,
		Name:     info.Name,
		ParentId: info.ParentId,
	}
	err := dao.Classify.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if classifyInfo.Id == 0 {
			//判断名称是否存在
			count, err := dao.Classify.Ctx(ctx).Where(do.Classify{Name: info.Name}).Count()
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "创建分类失败", dao.Classify.Table())
			}
			if count > 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "分类名称已存在"), "", dao.Classify.Table())
			}
			classifyInfo.Id = idgen.NextId()
			classifyInfo.CreatedAt = gtime.Now()

			// 清楚缓存
			_, err = dao.Classify.Ctx(ctx).Insert(classifyInfo)
			if err != nil {
				return err
			}
		} else {
			_, err := dao.Classify.Ctx(ctx).OmitEmptyData().Where(do.Classify{Id: classifyInfo.Id}).Update(do.Classify{
				Name:      classifyInfo.Name,
				ParentId:  classifyInfo.ParentId,
				UpdatedAt: classifyInfo.UpdatedAt,
			})
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新分类信息失败"), "", dao.Classify.Table())
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dao.Classify.Ctx(ctx).Where(do.Classify{Id: classifyInfo.Id}).Scan(&classifyInfo)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询分类ID失败", dao.Classify.Table())
	}

	return &classifyInfo, nil
}

func (s *sClassify) QueryClassify(ctx context.Context, id int64) (*entity.Classify, error) {
	result, err := daoctl.GetByIdWithError[entity.Classify](dao.Classify.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询分类失败", dao.Classify.Table())
	}
	return result, nil
}

func (s *sClassify) DeleteClassify(ctx context.Context, id int64) (bool, error) {
	info := &entity.Classify{}
	err := dao.Classify.Ctx(ctx).Where(do.Classify{Id: id}).Scan(&info)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除分类失败", dao.Classify.Table())
	}
	if info.Id == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除分类的ID不存在"), "", dao.Classify.Table())
	}

	//判断该分类下是否还有商品，有商品不能删除
	count, err := dao.Commodities.Ctx(ctx).Count(do.Commodities{ClassifyId: id})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除分类失败", dao.Classify.Table())
	}
	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该分类存在商品，不能删除"), "", dao.Classify.Table())
	}

	err = dao.Classify.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Classify.Ctx(ctx).Delete(do.Classify{Id: id})
		return nil
	})

	return err == nil, err
}
