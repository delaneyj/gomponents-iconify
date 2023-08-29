package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panflix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.368 43.626v-10.25a14.375 14.375 0 0 0 0-28.75H11.79a4.442 4.442 0 0 0-4.44 4.442v16.401c0 9.936 9.003 18.157 19.019 18.157Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.368 33.376v-5.961a8.36 8.36 0 0 0 0-16.72h-8.452a2.58 2.58 0 0 0-2.576 2.583v9.539c0 5.778 5.22 10.559 11.028 10.559Z"/>`),
		g.Group(children),
	)
}