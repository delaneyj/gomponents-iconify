package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#C60C30" d="M.5.5h300v226H.5z"/><path fill="#FFF" d="M96.791.5h32.354v226H96.791z"/><path fill="#FFF" d="M300.5 97.323v32.354H.5V97.323z"/></g>`),
		g.Group(children),
	)
}