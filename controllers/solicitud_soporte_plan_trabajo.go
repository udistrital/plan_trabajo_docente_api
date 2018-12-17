package controllers

import (
	"encoding/json"
	"errors"
	"plan_trabajo_docente_api/models"
	"strconv"
	"strings"
	"fmt"
	"github.com/astaxie/beego"
	//"reflect"
)

// SolicitudSoportePlanTrabajoController operations for SolicitudSoportePlanTrabajo
type SolicitudSoportePlanTrabajoController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudSoportePlanTrabajoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ObtenerCedulasSolicitudes", c.ObtenerCedulasSolicitudes)
}

// Post ...
// @Title Post
// @Description create SolicitudSoportePlanTrabajo
// @Param	body		body 	models.SolicitudSoportePlanTrabajo	true		"body for SolicitudSoportePlanTrabajo content"
// @Success 201 {int} models.SolicitudSoportePlanTrabajo
// @Failure 403 body is empty
// @router / [post]
func (c *SolicitudSoportePlanTrabajoController) Post() {
	var v models.SolicitudSoportePlanTrabajo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSolicitudSoportePlanTrabajo(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SolicitudSoportePlanTrabajo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SolicitudSoportePlanTrabajo
// @Failure 404 :no content found
// @router /:id [get]
func (c *SolicitudSoportePlanTrabajoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSolicitudSoportePlanTrabajoById(id)
	if err != nil {
		//c.Data["json"] = err.Error()
		beego.Error(err)
		c.Abort("404")

	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SolicitudSoportePlanTrabajo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	orCondition	query	string	false	"Recibe el campo y los valores para hacer la operacion OR con esos valores. e.g.col1:v1,col1:v2,col2:v3"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SolicitudSoportePlanTrabajo
// @Failure 403
// @router / [get]
func (c *SolicitudSoportePlanTrabajoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var orCondition = make(map[string][]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	fmt.Println(c.GetString("query"))
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// orCondition: k:v,k:v
	if v := c.GetString("orCondition"); v != "" {
		for _, cond := range strings.Split(v, ",") {
		 	kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid orCondition key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
		 	orCondition[k] = append(orCondition[k], v)
		}
	}

	l, err := models.GetAllSolicitudSoportePlanTrabajo(query, fields, sortby, order, orCondition,offset, limit)
	if err != nil {
		//c.Data["json"] = err.Error()
		beego.Error(err)
		c.Abort("404")
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SolicitudSoportePlanTrabajo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SolicitudSoportePlanTrabajo	true		"body for SolicitudSoportePlanTrabajo content"
// @Success 200 {object} models.SolicitudSoportePlanTrabajo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SolicitudSoportePlanTrabajoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SolicitudSoportePlanTrabajo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSolicitudSoportePlanTrabajoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SolicitudSoportePlanTrabajo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SolicitudSoportePlanTrabajoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSolicitudSoportePlanTrabajo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// ObtenerCedulasSolicitudes ...
// @Title Obtener Cedulas Solicitudes
// @Description get persona.id que pertenecen a SolicitudSoportePlanTrabajo
// @Param	iddependencia	query 	string	true	"id de la dependencia"
// @Param	estados	query	string	true	"Estados que se necesitan recuperar. e.g:estado:1,estado:2"
// @Param	anio	query	string	true	"anio de la solicitud"
// @Param	periodo	query	string	true	"semestre de la solicitud"
// @Success 200 {object} models.SolicitudSoportePlanTrabajo
// @Failure 403
// @router /obtener_cedulas [get]
func (c *SolicitudSoportePlanTrabajoController) ObtenerCedulasSolicitudes() {

	dependencia := c.GetString("iddependencia")
	anio := c.GetString("anio")
	periodo := c.GetString("periodo")

	var query = make(map[string]string)
	var orCondition = make(map[string][]string)
	query["organizacion"] = dependencia
	query["anio"] = anio
	query["periodo"] = periodo

	
	// orCondition: k:v,k:v
	if v := c.GetString("orCondition"); v != "" {
		for _, cond := range strings.Split(v, ",") {
		 	kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid orCondition key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
		 	orCondition[k] = append(orCondition[k], v)
		}
	}	


	listaCedulas,err := models.ObtenerCedulasSolicitudes(query, orCondition)
	if(err == nil ){
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = listaCedulas
	} else {
		beego.Error(err)
		c.Abort("404")
	}
	c.ServeJSON()
}

