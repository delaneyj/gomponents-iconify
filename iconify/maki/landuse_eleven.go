package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanduseEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1 4.749L.995 2.057a.251.251 0 0 1 .1-.2l2.249-1.8a.251.251 0 0 1 .313-.005l2.249 1.8a.251.251 0 0 1 .094.2v2.7A.251.251 0 0 1 5.751 5h-1.5A.251.251 0 0 1 4 4.749V3H3v1.752A.251.251 0 0 1 2.746 5h-1.5A.247.247 0 0 1 1 4.749zm4.753 2.6a.248.248 0 0 0-.173.072L4 9V6.5a.5.5 0 0 0-1 0V9H2V6.5a.5.5 0 0 0-1 0v4.25a.25.25 0 0 0 .25.25h4.5a.249.249 0 0 0 .25-.248V7.6a.25.25 0 0 0-.247-.253zM11 3.253v6.5a.247.247 0 0 1-.247.247H7.247A.247.247 0 0 1 7 9.753v-6.5A.252.252 0 0 1 7.252 3H8v-.752A.248.248 0 0 1 8.248 2h1.506a.246.246 0 0 1 .246.246V3h.747a.253.253 0 0 1 .253.253zM10 8H8v1h2zm0-2H8v1h2zm0-2H8v1h2z" fill="currentColor"/>`),
		g.Group(children),
	)
}