package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#002A8F" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="M.5 30.5h300v30H.5zm0 60h300v30H.5z"/><path fill="#CF142B" d="M.5.5v150l129.902-75z"/><path fill="#FFF" d="M43.801 53L38.75 68.551H22.402l13.219 9.609l-5.051 15.539L43.8 84.09l13.23 9.609l-5.051-15.539l13.219-9.609H48.85z"/></g>`),
		g.Group(children),
	)
}