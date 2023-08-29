package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D82727" d="M0 0h300v150H0z"/><path fill="#FFC71F" d="m0 0l150.226 75L0 150z"/><path fill="#000" d="m0 0l74.933 75L0 150z"/><path fill="#FFF" d="m33.499 65.953l15.823-6.594l-8.698 14.932l10.969 12.943l-16.854-4.09L25.703 98L24.31 80.57L7.433 76.78l15.948-6.606l-1.351-17.276z"/></g>`),
		g.Group(children),
	)
}