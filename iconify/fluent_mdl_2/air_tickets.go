package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirTickets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 640v1152H0V640h440q1 0 33-10t87-29t126-42t154-51t169-57t172-58t161-54t139-46t103-35t56-19l134 401h274zm-1206 0h797l-79-239l-718 239zm1078 128H128v896h1792V768zm-256 128v128h-128V896h128zm0 256v128h-128v-128h128zm0 384h-128v-128h128v128zm128-384h-128v-128h128v128zm0 256h-128v-128h128v128zM469 1280l-85-256h128l64 128h277l-85-256h128l128 256h256q26 0 45 19t19 45q0 26-19 45t-45 19h-256l-128 256H768l85-256H469z"/>`),
		g.Group(children),
	)
}