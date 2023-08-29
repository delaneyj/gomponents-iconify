package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func So(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#4189DD" d="M0 0h300v200H0z"/><path fill="#FFF" d="m104.208 85.122l28.301 20.561l-10.81 33.27L150 118.391l28.301 20.562l-10.81-33.27l28.301-20.561H160.81L150 51.852l-10.81 33.27z"/></g>`),
		g.Group(children),
	)
}