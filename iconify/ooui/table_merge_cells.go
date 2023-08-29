package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMergeCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path id="oouiTableMergeCells0" fill="currentColor" d="M9 10L4 6v3H0v2h4v3zm-7 3H0v5h8v-2H2zM0 2v5h2V4h6V2z"/><use href="#oouiTableMergeCells0" transform="matrix(-1 0 0 1 20 0)"/>`),
		g.Group(children),
	)
}