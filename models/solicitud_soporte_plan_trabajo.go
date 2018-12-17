package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SolicitudSoportePlanTrabajo struct {
	Id            int                       `orm:"column(id);pk;auto"`
	Estado        *EstadoSoportePlanTrabajo `orm:"column(estado);rel(fk)"`
	Persona       int                       `orm:"column(persona)"`
	Organizacion  int                       `orm:"column(organizacion)"`
	Documento     int                       `orm:"column(documento)"`
	Observaciones string                    `orm:"column(observaciones)"`
	Anio          int                       `orm:"column(anio)"`
	Periodo       int                       `orm:"column(periodo)"`
	Actividad     int                       `orm:"column(actividad)"`
}

func (t *SolicitudSoportePlanTrabajo) TableName() string {
	return "solicitud_soporte_plan_trabajo"
}

func init() {
	orm.RegisterModel(new(SolicitudSoportePlanTrabajo))
}

// AddSolicitudSoportePlanTrabajo insert a new SolicitudSoportePlanTrabajo into database and returns
// last inserted Id on success.
func AddSolicitudSoportePlanTrabajo(m *SolicitudSoportePlanTrabajo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudSoportePlanTrabajoById retrieves SolicitudSoportePlanTrabajo by Id. Returns error if
// Id doesn't exist
func GetSolicitudSoportePlanTrabajoById(id int) (v *SolicitudSoportePlanTrabajo, err error) {
	o := orm.NewOrm()
	v = &SolicitudSoportePlanTrabajo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudSoportePlanTrabajo retrieves all SolicitudSoportePlanTrabajo matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudSoportePlanTrabajo(query map[string]string, fields []string, sortby []string, order []string,
	orCondition map[string][]string,offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudSoportePlanTrabajo)).RelatedSel()

	// se añade soporte para condicion OR
	if(len(orCondition) > 0){
		cond := orm.NewCondition()
		cond1 := orm.NewCondition()
		for k, v := range orCondition {
			for _,valor := range v{
				cond = cond.Or(k, valor)
			}	
		}
		cond1 = cond1.AndCond(cond)

		qs = qs.SetCond(cond1)
	}
	

	// query k=v
	for k, v := range query {

		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SolicitudSoportePlanTrabajo
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSolicitudSoportePlanTrabajo updates SolicitudSoportePlanTrabajo by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudSoportePlanTrabajoById(m *SolicitudSoportePlanTrabajo) (err error) {
	o := orm.NewOrm()
	v := SolicitudSoportePlanTrabajo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudSoportePlanTrabajo deletes SolicitudSoportePlanTrabajo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudSoportePlanTrabajo(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudSoportePlanTrabajo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudSoportePlanTrabajo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//trae las cedulas de las solicitudes
//podria consultar desde acá personas?
func ObtenerCedulasSolicitudes(query map[string]string, orCondition map[string][]string) ([]orm.Params ,error) {
	o := orm.NewOrm()

	fmt.Println(query)
	fmt.Println(orCondition)

	var maps []orm.Params
	qs := o.QueryTable(new(SolicitudSoportePlanTrabajo)).Distinct().Limit(-1)
	// query k=v
	for k, v := range query {

		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	// se añade soporte para condicion OR
	cond := orm.NewCondition()
	for k, v := range orCondition {
		for _,valor := range v{
			cond = cond.Or(k, valor)
		}	
	}

	cond = cond.AndCond(cond)

	fmt.Println("_____________________")
	fmt.Println(cond)
	fmt.Println("_____________________")

	qs = qs.SetCond(cond)

	_, err := qs.Values(&maps, "Persona")
	
	fmt.Println("_________________________-")
	fmt.Println(maps)
	fmt.Println(len(maps))
	fmt.Println("_________________________-")

	// var temp []int
	// _, err := o.Raw("SELECT DISTINCT s.persona FROM academica.solicitud_soporte_plan_Trabajo s WHERE s.estado = ? and s.organizacion = ? and s.anio = ? and s.periodo = ? ",estado, dependencia ,anio ,periodo).QueryRows(&temp)

	if len(maps) == 0{
		err = orm.ErrNoRows
	}

	if err == nil {
		fmt.Println("Consulta exitosa")
		//fmt.Println(temp)
		return maps, nil
	}
	return maps, err
}