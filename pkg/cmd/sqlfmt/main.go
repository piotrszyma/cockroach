// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// smithtest is a tool to execute sqlsmith tests on cockroach demo
// instances. Failures are tracked, de-duplicated, reduced. Issues are
// prefilled for GitHub.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cockroachdb/errors"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// TODO: Read flags.
var (
	flagLen        int
	flagUseSpaces  bool
	flagTabWidth   int
	flagNoSimplify bool
	flagAlign      bool
	flagPath       string
)

// Goal:
// sqlfmt --format ./dir
// sqlfmt --format ./file.sql
type SqlfmtCtx struct {
	len        int
	useSpaces  bool
	tabWidth   int
	noSimplify bool
	align      bool
	formatPath string
}

func runSQLFmt(sqlfmtCtx SqlfmtCtx) error {
	if sqlfmtCtx.len < 1 {
		return errors.Errorf("line length must be > 0: %d", sqlfmtCtx.len)
	}
	if sqlfmtCtx.tabWidth < 1 {
		return errors.Errorf("tab width must be > 0: %d", sqlfmtCtx.tabWidth)
	}

	var sl parser.Statements
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	sl, err = parser.Parse(string(in))
	if err != nil {
		return err
	}

	cfg := tree.DefaultPrettyCfg()
	cfg.UseTabs = !sqlfmtCtx.useSpaces
	cfg.LineWidth = sqlfmtCtx.len
	cfg.TabWidth = sqlfmtCtx.tabWidth
	cfg.Simplify = !sqlfmtCtx.noSimplify
	cfg.Align = tree.PrettyNoAlign
	cfg.JSONFmt = true
	if sqlfmtCtx.align {
		cfg.Align = tree.PrettyAlignAndDeindent
	}

	for i := range sl {
		fmt.Print(cfg.Pretty(sl[i].AST))
		if len(sl) > 1 {
			fmt.Print(";")
		}
		fmt.Println()
	}
	return nil
}

func main() {
	flag.IntVar(&flagLen, "len", 4, "len")
	flag.BoolVar(&flagUseSpaces, "use-spaces", true, "use spaces")
	flag.IntVar(&flagTabWidth, "tab-width", 4, "tab width")
	flag.BoolVar(&flagNoSimplify, "no-simplify", false, "no simplify")
	flag.BoolVar(&flagAlign, "align", true, "align")

	flag.Parse()

	if flag.NArg() != 1 { // Expect one arg.
		flag.Usage()
		os.Exit(1)
	}

	formatPath := flag.Arg(1)

	runSQLFmt(SqlfmtCtx{
		len:        flagLen,
		useSpaces:  flagUseSpaces,
		tabWidth:   flagTabWidth,
		noSimplify: flagNoSimplify,
		align:      flagAlign,
		formatPath: formatPath,
	})
}
