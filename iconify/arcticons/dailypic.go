package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailypic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9.5h-1a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-28a2 2 0 0 0-2-2h-1m-6 0h-19m22 3v-6m-25 6v-6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.09 31.288A6.986 6.986 0 0 0 30.985 25.5m-3.079-5.792a6.986 6.986 0 0 0-10.89 5.792"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.906 25.5h-2.625l3.61-3.609L34.5 25.5h0h-4.594zm-11.812.001h2.625l-3.61 3.609l-3.609-3.609h4.594z"/>`),
		g.Group(children),
	)
}