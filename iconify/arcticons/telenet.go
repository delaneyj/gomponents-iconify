package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telenet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/><ellipse cx="19.062" cy="18.783" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.17" ry="2.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.274 20.324v-2.75a11.803 11.803 0 0 1 6.183.155v2.73a11.401 11.401 0 0 0-6.183-.135ZM33 26.637v1.876a12.84 12.84 0 0 1-18-.138v-1.876s9.537 6.665 18 .138Z"/>`),
		g.Group(children),
	)
}