package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zulip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 16c0 8.839-7.161 16-16 16S0 24.839 0 16S7.161 0 16 0s16 7.161 16 16zM21.292 8H9.36l1.333 2.708h8.078l-9.411 10l1.333 3.276h11.948l-1.333-2.693h-8.094l9.411-10l-1.333-3.276z"/>`),
		g.Group(children),
	)
}