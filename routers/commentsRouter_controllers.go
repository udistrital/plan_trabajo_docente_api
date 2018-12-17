package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:EstadoSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"] = append(beego.GlobalControllerRouter["plan_trabajo_docente_api/controllers:SolicitudSoportePlanTrabajoController"],
		beego.ControllerComments{
			Method: "ObtenerCedulasSolicitudes",
			Router: `/obtener_cedulas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
