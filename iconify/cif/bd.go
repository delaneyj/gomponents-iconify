package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#006A4E" d="M.5.5h300v180H.5z"/><circle cx="135.5" cy="90.5" r="60" fill="#F42A41"/></g>`),
		g.Group(children),
	)
}