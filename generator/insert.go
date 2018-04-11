package generator

import (
	"bytes"

	"github.com/arstd/light/goparser"
	"github.com/arstd/light/sqlparser"
)

func writeInsert(buf *bytes.Buffer, m *goparser.Method, stmt *sqlparser.Statement) {
	wln := func(s string) { buf.WriteString(s + "\n") }

	wln("var buf bytes.Buffer")
	wln("var args []interface{}")

	for _, f := range stmt.Fragments {
		writeFragment(buf, m, f)
	}

	wln("query := buf.String()")
	if m.Store.Log {
		wln("log.Debug(query)")
		wln("log.Debug(args...)")
	}

	wln(`ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		res, err := exec.ExecContext(ctx, query, args...)`)
	wln("if err != nil {")
	if m.Store.Log {
		wln("log.Error(query)")
		wln("log.Error(args...)")
		wln("log.Error(err)")
	}
	wln("return 0, err")
	wln("}")
	wln("return res.LastInsertId()")
}
