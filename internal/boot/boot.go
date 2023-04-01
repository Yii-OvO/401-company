package boot

import (
	"401-company/controller"
	"401-company/internal/consts"
	"401-company/model/enum"
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/SupenBysz/gf-admin-company-modules/example/router"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"
)

func init() {
	permissionTree := InitPermission()
	consts.PermissionTree = permissionTree
}

// InitPermission 初始化生成自定义权限树
func InitPermission() []*sys_model.SysPermissionTree {
	sys_consts.Global.PermissionTree = []*sys_model.SysPermissionTree{
		// 分类管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         6413852638249029,
				Name:       "分类管理",
				Identifier: "Classify",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看分类，查看某个分类
				enum.Classify.PermissionType.Query,
				// 更新分类信息，更新某个分类信息
				enum.Classify.PermissionType.Update,
				// 删除分类，删除某个分类
				enum.Classify.PermissionType.Delete,
				// 创建分类，创建一个新分类
				enum.Classify.PermissionType.Create,
			},
		},
		// 商品管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         6413877039857733,
				Name:       "商品管理",
				Identifier: "Commodity",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看商品，查看某个商品
				enum.Commodity.PermissionType.Query,
				// 更新商品信息，更新某个商品信息
				enum.Commodity.PermissionType.Update,
				// 删除商品，删除某个商品
				enum.Commodity.PermissionType.Delete,
				// 创建商品，创建一个新商品
				enum.Commodity.PermissionType.Create,
			},
		},
	}
	return sys_consts.Global.PermissionTree
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			{
				// OpenApi自定义信息
				oai.Info.Title = `API Reference`
				oai.Config.CommonResponse = api_v1.JsonRes{}
				oai.Config.CommonResponseDataField = `Data`
			}

			{
				// 静态目录设置
				uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
				if uploadPath == "" {
					g.Log().Fatal(ctx, "文件上传配置路径不能为空!")
				}
				if !gfile.Exists(uploadPath) {
					_ = gfile.Mkdir(uploadPath)
				}
				// 上传目录添加至静态资源
				s.AddStaticPath("/upload", uploadPath)
			}

			{
				// HOOK, 开发阶段禁止浏览器缓存,方便调试
				if gmode.IsDevelop() {
					s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
						r.Response.Header().Set("Cache-Control", "no-store")
					})
				}
			}

			{
				// ImportPermissionTree 导入权限结构
				sys_service.SysPermission().ImportPermissionTree(ctx, consts.PermissionTree, nil)

				// 导入财务服务权限结构 (可选)
				sys_service.SysPermission().ImportPermissionTree(ctx, consts.FinancialPermissionTree, nil)

				// CASBIN 初始化
				sys_service.Casbin().Enforcer()

			}

			// 初始化路由
			apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
			s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
				// 注册中间件
				group.Middleware(
					// sys_service.Middleware().Casbin,

					sys_service.Middleware().CTX,
					sys_service.Middleware().ResponseHandler,
				)

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 鉴权：登录，注册，找回密码等
					group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
					// 图型验证码、短信验证码、地区
					group.Group("/common", func(group *ghttp.RouterGroup) {
						group.Bind(
							// 图型验证码
							sys_controller.Captcha,
							// 短信验证码
							sys_controller.SysSms,
							// 地区
							sys_controller.SysArea,
						)
					})

				})

				// 权限路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						// sys_service.Middleware().CTX,
						sys_service.Middleware().Auth,
					)

					// 注册公司模块路由 （包含：公司、团队、员工）
					router.ModulesGroup(consts.Global.Modules, group)

					// 注册财务模块路由 (可选)
					router.FinancialGroup(consts.Global.Modules, group)

					// 分类
					group.Group("/classify", func(group *ghttp.RouterGroup) { group.Bind(controller.Classify) })
					// 商品
					group.Group("/commodity", func(group *ghttp.RouterGroup) { group.Bind(controller.Commodity) })
				})
			})
			s.Run()
			return nil
		},
	}
)
