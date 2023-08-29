package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024q-31 0-56-32L128 263v697q0 26-18.5 45T64 1024t-45.5-19T0 960V65q0-27 18.5-46T64 0q35 0 55 32l521 730V64q0-26 19-45t45.5-19t45 19T768 64v896q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}