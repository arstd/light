// DO NOT EDIT THIS FILE!
// It is generated by `light` tool from source `model.go`.

package mapper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/arstd/light/example/domain"
	"github.com/arstd/light/example/enum"
	"github.com/arstd/log"
	"github.com/lib/pq"
)

type ModelMapperImpl struct{}

// insert into models(name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32) values (${m.Name}, ${m.Flag}, ${m.Score}, ${m.Map}, ${m.Time}, ${m.Array}, ${m.Slice}, ${m.Status}, ${m.Pointer}, ${m.StructSlice}, ${m.Uint32}) returning id
func (*ModelMapperImpl) Insert(m *domain.Model, xtx ...*sql.Tx) (err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// insert into models(name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32) values (${m.Name}, ${m.Flag}, ${m.Score}, ${m.Map}, ${m.Time}, ${m.Array}, ${m.Slice}, ${m.Status}, ${m.Pointer}, ${m.StructSlice}, ${m.Uint32}) returning id
	xbuf.WriteString(`insert into models(name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32) values (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) returning id `)
	zmMap, _ := json.Marshal(m.Map)
	zmPointer, _ := json.Marshal(m.Pointer)
	zmStructSlice, _ := json.Marshal(m.StructSlice)
	xargs = append(xargs, m.Name, m.Flag, m.Score, zmMap, m.Time, pq.Array(m.Array), pq.Array(m.Slice), m.Status, zmPointer, zmStructSlice, time.Unix(int64(m.Uint32), 0))

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	xdest := []interface{}{&m.Id}
	if len(xtx) > 0 {
		err = xtx[0].QueryRow(xquery, xargs...).Scan(xdest...)
	} else {
		err = db.QueryRow(xquery, xargs...).Scan(xdest...)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
	}
	return
}

// insert into models(uint32, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice) values [{ i, m := range ms | , } (${i}+888, ${m.Name}, ${m.Flag}, ${m.Score}, ${m.Map}, ${m.Time}, ${m.Array}, ${m.Slice}, ${m.Status}, ${m.Pointer}, ${m.StructSlice}) ]
func (*ModelMapperImpl) BatchInsert(ms []*domain.Model, xtx ...*sql.Tx) (xa int64, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// insert into models(uint32, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice) values
	xbuf.WriteString(`insert into models(uint32, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice) values `)
	// (${i}+888, ${m.Name}, ${m.Flag}, ${m.Score}, ${m.Map}, ${m.Time}, ${m.Array}, ${m.Slice}, ${m.Status}, ${m.Pointer}, ${m.StructSlice})
	for i, m := range ms {
		if i != 0 {
			xbuf.WriteString(", ")
		}
		xbuf.WriteString(`(%s+888, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) `)
		zmMap, _ := json.Marshal(m.Map)
		zmPointer, _ := json.Marshal(m.Pointer)
		zmStructSlice, _ := json.Marshal(m.StructSlice)
		xargs = append(xargs, i, m.Name, m.Flag, m.Score, zmMap, m.Time, pq.Array(m.Array), pq.Array(m.Slice), m.Status, zmPointer, zmStructSlice)
	}

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	var xres sql.Result
	if len(xtx) > 0 {
		xres, err = xtx[0].Exec(xquery, xargs...)
	} else {
		xres, err = db.Exec(xquery, xargs...)
	}
	if err != nil {
		log.Error(xquery)
		log.Error(xargs...)
		log.Error(err)
	}
	return xres.RowsAffected()
}

// select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where id=${id}
func (*ModelMapperImpl) Get(id int, xtx ...*sql.Tx) (xobj *domain.Model, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where id=${id}
	xbuf.WriteString(`select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where id=%s `)
	xargs = append(xargs, id)

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	xobj = &domain.Model{}
	var xobjzMap []byte
	var xobjzTime pq.NullTime
	var xobjzPointer []byte
	var xobjzStructSlice []byte
	var xobjzUint32 pq.NullTime
	xdest := []interface{}{&xobj.Id, &xobj.Name, &xobj.Flag, &xobj.Score, &xobjzMap,
		&xobjzTime, pq.Array(&xobj.Array), pq.Array(&xobj.Slice), &xobj.Status, &xobjzPointer, &xobjzStructSlice,
		&xobjzUint32}
	if len(xtx) > 0 {
		err = xtx[0].QueryRow(xquery, xargs...).Scan(xdest...)
	} else {
		err = db.QueryRow(xquery, xargs...).Scan(xdest...)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
	}
	xobj.Map = map[string]interface{}{}
	json.Unmarshal(xobjzMap, xobj.Map)
	xobj.Time = xobjzTime.Time
	xobj.Pointer = &domain.Model{}
	json.Unmarshal(xobjzPointer, &xobj.Pointer)
	xobj.StructSlice = []*domain.Model{}
	json.Unmarshal(xobjzStructSlice, &xobj.StructSlice)
	xobj.Uint32 = uint32(xobjzUint32.Time.Unix())
	return
}

// update models set name=${m.Name}, flag=${m.Flag}, score=${m.Score}, map=${m.Map}, time=${m.Time}, slice=${m.Slice}, status=${m.Status}, pointer=${m.Pointer}, struct_slice=${m.StructSlice}, uint32=${m.Uint32} where id=${m.Id}
func (*ModelMapperImpl) Update(m *domain.Model, xtx ...*sql.Tx) (xa int64, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// update models set name=${m.Name}, flag=${m.Flag}, score=${m.Score}, map=${m.Map}, time=${m.Time}, slice=${m.Slice}, status=${m.Status}, pointer=${m.Pointer}, struct_slice=${m.StructSlice}, uint32=${m.Uint32} where id=${m.Id}
	xbuf.WriteString(`update models set name=%s, flag=%s, score=%s, map=%s, time=%s, slice=%s, status=%s, pointer=%s, struct_slice=%s, uint32=%s where id=%s `)
	zmMap, _ := json.Marshal(m.Map)
	zmPointer, _ := json.Marshal(m.Pointer)
	zmStructSlice, _ := json.Marshal(m.StructSlice)
	xargs = append(xargs, m.Name, m.Flag, m.Score, zmMap, m.Time, pq.Array(m.Slice), m.Status, zmPointer, zmStructSlice, time.Unix(int64(m.Uint32), 0), m.Id)

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	var xres sql.Result
	if len(xtx) > 0 {
		xres, err = xtx[0].Exec(xquery, xargs...)
	} else {
		xres, err = db.Exec(xquery, xargs...)
	}
	if err != nil {
		log.Error(xquery)
		log.Error(xargs...)
		log.Error(err)
	}
	return xres.RowsAffected()
}

// delete from models where id=${id}
func (*ModelMapperImpl) Delete(id int, xtx ...*sql.Tx) (xa int64, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// delete from models where id=${id}
	xbuf.WriteString(`delete from models where id=%s `)
	xargs = append(xargs, id)

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	var xres sql.Result
	if len(xtx) > 0 {
		xres, err = xtx[0].Exec(xquery, xargs...)
	} else {
		xres, err = db.Exec(xquery, xargs...)
	}
	if err != nil {
		log.Error(xquery)
		log.Error(xargs...)
		log.Error(err)
	}
	return xres.RowsAffected()
}

// select count(*) from models where name like ${m.Name} [{ m.Flag } and flag=${m.Flag} ] [{ len(m.Array) != 0 } and xarray && array[ [{range m.Array}] ] ] [{ len(ss) != 0 } and status in ( [{range ss}] ) ] [{ len(m.Slice) != 0 } and slice && ${m.Slice} ]
func (*ModelMapperImpl) Count(m *domain.Model, ss []enum.Status, xtx ...*sql.Tx) (xcnt int64, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// select count(*) from models where name like ${m.Name}
	xbuf.WriteString(`select count(*) from models where name like %s `)
	xargs = append(xargs, m.Name)
	// and flag=${m.Flag}
	if m.Flag {
		xbuf.WriteString(`and flag=%s `)
		xargs = append(xargs, m.Flag)
	}
	// and xarray && array[ [{range m.Array}] ]
	if len(m.Array) != 0 {
		// and xarray && array[
		xbuf.WriteString(`and xarray && array[ `)
		// ${v}
		for i, v := range m.Array {
			if i != 0 {
				xbuf.WriteString(", ")
			}
			xbuf.WriteString(`%s `)
			xargs = append(xargs, v)
		}
		// ]
		xbuf.WriteString(`] `)
	}
	// and status in ( [{range ss}] )
	if len(ss) != 0 {
		// and status in (
		xbuf.WriteString(`and status in ( `)
		// ${v}
		for i, v := range ss {
			if i != 0 {
				xbuf.WriteString(", ")
			}
			xbuf.WriteString(`%s `)
			xargs = append(xargs, v)
		}
		// )
		xbuf.WriteString(`) `)
	}
	// and slice && ${m.Slice}
	if len(m.Slice) != 0 {
		xbuf.WriteString(`and slice && %s `)
		xargs = append(xargs, pq.Array(m.Slice))
	}

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	if len(xtx) > 0 {
		err = xtx[0].QueryRow(xquery, xargs...).Scan(&xcnt)
	} else {
		err = db.QueryRow(xquery, xargs...).Scan(&xcnt)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
	}
	return
}

// select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where name like ${m.Name} [ [ and status in ( [{range ss}] ) ] and flag=${m.Flag} [{ !from.IsZero() && to.IsZero() } and time >= ${from} ] ] [ and time between ${from} and ${to} ] [{ !from.IsZero() && to.IsZero() } and time >= ${from} ] [{ from.IsZero() && !to.IsZero() } and time <= ${to} ] [ and xarray && array[ [{range m.Array}] ] ] [ and slice && ${m.Slice} ] order by id offset ${offset} limit ${limit}
func (*ModelMapperImpl) List(m *domain.Model, ss []enum.Status, from time.Time, to time.Time, offset int, limit int, xtx ...*sql.Tx) (xdata []*domain.Model, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where name like ${m.Name}
	xbuf.WriteString(`select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where name like %s `)
	xargs = append(xargs, m.Name)
	// [ and status in ( [{range ss}] ) ] and flag=${m.Flag} [{ !from.IsZero() && to.IsZero() } and time >= ${from} ]
	if len(ss) != 0 && m.Flag && !from.IsZero() && to.IsZero() {
		// and status in ( [{range ss}] )
		if len(ss) != 0 {
			// and status in (
			xbuf.WriteString(`and status in ( `)
			// ${v}
			for i, v := range ss {
				if i != 0 {
					xbuf.WriteString(", ")
				}
				xbuf.WriteString(`%s `)
				xargs = append(xargs, v)
			}
			// )
			xbuf.WriteString(`) `)
		}
		// and flag=${m.Flag}
		if m.Flag {
			xbuf.WriteString(`and flag=%s `)
			xargs = append(xargs, m.Flag)
		}
		// and time >= ${from}
		if !from.IsZero() && to.IsZero() {
			xbuf.WriteString(`and time >= %s `)
			xargs = append(xargs, from)
		}
	}
	// and time between ${from} and ${to}
	if !from.IsZero() && !to.IsZero() {
		xbuf.WriteString(`and time between %s and %s `)
		xargs = append(xargs, from, to)
	}
	// and time >= ${from}
	if !from.IsZero() && to.IsZero() {
		xbuf.WriteString(`and time >= %s `)
		xargs = append(xargs, from)
	}
	// and time <= ${to}
	if from.IsZero() && !to.IsZero() {
		xbuf.WriteString(`and time <= %s `)
		xargs = append(xargs, to)
	}
	// and xarray && array[ [{range m.Array}] ]
	if len(m.Array) != 0 {
		// and xarray && array[
		xbuf.WriteString(`and xarray && array[ `)
		// ${v}
		for i, v := range m.Array {
			if i != 0 {
				xbuf.WriteString(", ")
			}
			xbuf.WriteString(`%s `)
			xargs = append(xargs, v)
		}
		// ]
		xbuf.WriteString(`] `)
	}
	// and slice && ${m.Slice}
	if len(m.Slice) != 0 {
		xbuf.WriteString(`and slice && %s `)
		xargs = append(xargs, pq.Array(m.Slice))
	}
	// order by id offset ${offset} limit ${limit}
	xbuf.WriteString(`order by id offset %s limit %s `)
	xargs = append(xargs, offset, limit)

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)
	log.Debug(xquery)
	log.Debug(xargs...)

	var xrows *sql.Rows
	if len(xtx) > 0 {
		xrows, err = xtx[0].Query(xquery, xargs...)
	} else {
		xrows, err = db.Query(xquery, xargs...)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
		return
	}
	defer xrows.Close()

	xdata = []*domain.Model{}
	for xrows.Next() {
		xe := &domain.Model{}
		xdata = append(xdata, xe)
		var xezMap []byte
		var xezTime pq.NullTime
		var xezPointer []byte
		var xezStructSlice []byte
		var xezUint32 pq.NullTime
		xdest := []interface{}{&xe.Id, &xe.Name, &xe.Flag, &xe.Score, &xezMap,
			&xezTime, pq.Array(&xe.Array), pq.Array(&xe.Slice), &xe.Status, &xezPointer, &xezStructSlice,
			&xezUint32}
		err = xrows.Scan(xdest...)
		if err != nil {
			log.Error(err)
			return
		}
		xe.Map = map[string]interface{}{}
		json.Unmarshal(xezMap, xe.Map)
		xe.Time = xezTime.Time
		xe.Pointer = &domain.Model{}
		json.Unmarshal(xezPointer, &xe.Pointer)
		xe.StructSlice = []*domain.Model{}
		json.Unmarshal(xezStructSlice, &xe.StructSlice)
		xe.Uint32 = uint32(xezUint32.Time.Unix())
	}
	if err = xrows.Err(); err != nil {
		log.Error(err)
	}
	return
}

// select id, name, flag, score, map, time, slice, status, pointer, struct_slice from models where name like ${m.Name} [{ m.Flag != false } [{ len(ss) != 0 } and status in ( [{range ss}] ) ] and flag=${m.Flag} ] [{ len(m.Slice) != 0 } and slice && ${m.Slice} ] [ and time between ${from} and ${to} ] [{ !from.IsZero() && to.IsZero() } and time >= ${from} ] [{ from.IsZero() && !to.IsZero() } and time <= ${to} ] order by id offset ${offset} limit ${limit}
func (*ModelMapperImpl) Page(m *domain.Model, ss []enum.Status, from time.Time, to time.Time, offset int, limit int, xtx ...*sql.Tx) (xcnt int64, xdata []*domain.Model, err error) {
	var (
		xbuf  bytes.Buffer
		xargs []interface{}
	)
	// select id, name, flag, score, map, time, slice, status, pointer, struct_slice from models where name like ${m.Name}
	xbuf.WriteString(`select id, name, flag, score, map, time, slice, status, pointer, struct_slice from models where name like %s `)
	xargs = append(xargs, m.Name)
	// [{ len(ss) != 0 } and status in ( [{range ss}] ) ] and flag=${m.Flag}
	if m.Flag != false {
		// and status in ( [{range ss}] )
		if len(ss) != 0 {
			// and status in (
			xbuf.WriteString(`and status in ( `)
			// ${v}
			for i, v := range ss {
				if i != 0 {
					xbuf.WriteString(", ")
				}
				xbuf.WriteString(`%s `)
				xargs = append(xargs, v)
			}
			// )
			xbuf.WriteString(`) `)
		}
		// and flag=${m.Flag}
		xbuf.WriteString(`and flag=%s `)
		xargs = append(xargs, m.Flag)
	}
	// and slice && ${m.Slice}
	if len(m.Slice) != 0 {
		xbuf.WriteString(`and slice && %s `)
		xargs = append(xargs, pq.Array(m.Slice))
	}
	// and time between ${from} and ${to}
	if !from.IsZero() && !to.IsZero() {
		xbuf.WriteString(`and time between %s and %s `)
		xargs = append(xargs, from, to)
	}
	// and time >= ${from}
	if !from.IsZero() && to.IsZero() {
		xbuf.WriteString(`and time >= %s `)
		xargs = append(xargs, from)
	}
	// and time <= ${to}
	if from.IsZero() && !to.IsZero() {
		xbuf.WriteString(`and time <= %s `)
		xargs = append(xargs, to)
	}
	// order by id offset ${offset} limit ${limit}
	xbuf.WriteString(`order by id offset %s limit %s `)
	xargs = append(xargs, offset, limit)

	xholder := make([]interface{}, len(xargs))
	for i := range xargs {
		xholder[i] = fmt.Sprintf("$%d", i+1)
	}
	xquery := fmt.Sprintf(xbuf.String(), xholder...)

	xfindex := strings.LastIndex(xquery, " from ")
	xobindex := strings.LastIndex(xquery, "order by")
	xtquery := `select count(*)` + xquery[xfindex:xobindex]
	xdcnt := strings.Count(xquery[xobindex:], "$")
	xtargs := xargs[:len(xargs)-xdcnt]
	log.Debug(xtquery)
	log.Debug(xtargs...)

	if len(xtx) > 0 {
		err = xtx[0].QueryRow(xtquery, xtargs...).Scan(&xcnt)
	} else {
		err = db.QueryRow(xtquery, xtargs...).Scan(&xcnt)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
		return
	}
	if xcnt == 0 {
		return
	}

	log.Debug(xquery)
	log.Debug(xargs...)

	var xrows *sql.Rows

	if len(xtx) > 0 {
		xrows, err = xtx[0].Query(xquery, xargs...)
	} else {
		xrows, err = db.Query(xquery, xargs...)
	}
	if err != nil {
		log.Error(err)
		log.Error(xquery)
		log.Error(xargs...)
		return
	}
	defer xrows.Close()

	xdata = []*domain.Model{}
	for xrows.Next() {
		xe := &domain.Model{}
		xdata = append(xdata, xe)
		var xezMap []byte
		var xezTime pq.NullTime
		var xezPointer []byte
		var xezStructSlice []byte
		xdest := []interface{}{&xe.Id, &xe.Name, &xe.Flag, &xe.Score, &xezMap,
			&xezTime, pq.Array(&xe.Slice), &xe.Status, &xezPointer, &xezStructSlice}
		err = xrows.Scan(xdest...)
		if err != nil {
			log.Error(err)
			return
		}
		xe.Map = map[string]interface{}{}
		json.Unmarshal(xezMap, xe.Map)
		xe.Time = xezTime.Time
		xe.Pointer = &domain.Model{}
		json.Unmarshal(xezPointer, &xe.Pointer)
		xe.StructSlice = []*domain.Model{}
		json.Unmarshal(xezStructSlice, &xe.StructSlice)
	}
	if err = xrows.Err(); err != nil {
		log.Error(err)
	}
	return
}
