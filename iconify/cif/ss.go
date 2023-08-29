package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#078930" d="M.5 105.5h300v45H.5z"/><path fill="#FFF" d="M.5 45.5h300v60H.5z"/><path fill="#000" d="M.5.5h300v45H.5z"/><path fill="#DA121A" d="M.5 53h300v45H.5z"/><path fill="#0F47AF" d="m.5.5l129.904 75L.5 150.5z"/><path fill="#FCDD09" d="m19.801 75.5l43.417 14.107l-26.833-36.932v45.65l26.833-36.932z"/></g>`),
		g.Group(children),
	)
}