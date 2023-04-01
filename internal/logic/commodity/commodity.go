package commodity

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
)

type sCommodity struct{}

func init() {
	service.RegisterCommodity(New())
}

func New() *sCommodity {
	return &sCommodity{}
}

func (s *sCommodity) CreateCommodity(ctx context.Context, info model.Commodity) (*entity.Commodities, error) {
	info.Id = 0
	return s.SaveCommodity(ctx, info)
}

func (s *sCommodity) UpdateCommodity(ctx context.Context, info model.Commodity) (*entity.Commodities, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新商品信息ID参数错误"), "", dao.Commodities.Table())
	}
	return s.SaveCommodity(ctx, info)
}

func (s *sCommodity) SaveCommodity(ctx context.Context, info model.Commodity) (*entity.Commodities, error) {
	commodityInfo := entity.Commodities{
		Id:         info.Id,
		Name:       info.Name,
		ClassifyId: info.ClassifyId,
		Des:        info.Des,
	}
	err := dao.Commodities.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if commodityInfo.Id == 0 {
			//判断商品是否分类
			count, err := dao.Classify.Ctx(ctx).Count(do.Classify{
				Id: info.ClassifyId,
			})
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "创建商品失败", dao.Commodities.Table())
			}
			if count <= 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "商品未分类"), "", dao.Commodities.Table())
			}

			//判断名称是否存在
			count, err = dao.Commodities.Ctx(ctx).Where(do.Commodities{Name: info.Name}).Count()
			if count > 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "商品名称已存在"), "", dao.Commodities.Table())
			}
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "创建商品失败", dao.Commodities.Table())
			}

			commodityInfo.Id = idgen.NextId()
			commodityInfo.CreatedAt = gtime.Now()
			
			//清楚缓存
			_, err = dao.Commodities.Ctx(ctx).Insert(commodityInfo)
			if err != nil {
				return err
			}

		} else {
			//判断商品是否分类
			count, err := dao.Classify.Ctx(ctx).Count(do.Classify{
				Id: info.ClassifyId,
			})
			if count <= 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "商品未分类"), "", dao.Commodities.Table())
			}
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "创建商品失败", dao.Commodities.Table())
			}

			_, err = dao.Commodities.Ctx(ctx).OmitEmptyData().Where(do.Commodities{Id: commodityInfo.Id}).Update(do.Commodities{
				Name:       commodityInfo.Name,
				ClassifyId: commodityInfo.ClassifyId,
				Des:        commodityInfo.Des,
				UpdatedAt:  commodityInfo.UpdatedAt,
			})
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新商品信息失败"), "", dao.Commodities.Table())
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dao.Commodities.Ctx(ctx).Where(do.Commodities{Id: commodityInfo.Id}).Scan(&commodityInfo)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询商品ID失败", dao.Commodities.Table())
	}

	return &commodityInfo, err
}

func (s *sCommodity) QueryCommodity(ctx context.Context, id int64) (*entity.Commodities, error) {
	result, err := daoctl.GetByIdWithError[entity.Commodities](dao.Commodities.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询商品失败", dao.Commodities.Table())
	}
	return result, nil
}

func (s *sCommodity) DeleteCommodity(ctx context.Context, id int64) (bool, error) {
	info := &entity.Commodities{}
	err := dao.Commodities.Ctx(ctx).Where(do.Commodities{Id: id}).Scan(&info)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除商品失败", dao.Commodities.Table())
	}
	if info.Id == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除分类的ID不存在"), "", dao.Commodities.Table())
	}

	err = dao.Commodities.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Commodities.Ctx(ctx).Delete(do.Commodities{Id: id})
		return nil
	})

	return err == nil, err
}
