package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 0a18.955 18.955 0 0 0-18.955 18.957A18.955 18.955 0 0 0 47.5 37.747V97.5A2.5 2.5 0 0 0 50 100a2.5 2.5 0 0 0 2.5-2.5V37.746a18.955 18.955 0 0 0 16.455-18.789A18.955 18.955 0 0 0 50 0zm-4.832 8.549a5.947 5.947 0 0 1 5.947 5.947a5.947 5.947 0 0 1-5.947 5.947a5.947 5.947 0 0 1-5.947-5.947a5.947 5.947 0 0 1 5.947-5.947z" color="currentColor"/>`),
		g.Group(children),
	)
}