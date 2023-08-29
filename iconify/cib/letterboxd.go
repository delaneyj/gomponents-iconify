package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Letterboxd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11.052 22.339V9.599H8.729V6.401h8.438v3.198h-2.328v12.766h5.234v-3.49h3.781v6.724H8.729v-3.26zM0 16c0 8.839 7.161 16 16 16s16-7.161 16-16S24.839 0 16 0S0 7.161 0 16z"/>`),
		g.Group(children),
	)
}