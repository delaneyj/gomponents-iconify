package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M237 3q79 0 112 39q30 35 20 99q-23 146-175 146h-49q-8 0-14 5t-7 13l-17 106q-1 8-7 13t-14 5H13q-6 0-10-4.5T0 415L62 21q2-8 7.5-13T83 3h154zm18 144q4-29-8-43q-6-8-18-11.5t-21.5-4T180 88h-11q-11 0-12 11l-17 103h23q17 0 25.5-.5t22-3.5t21-8.5t14-16.5t9.5-26z"/>`),
		g.Group(children),
	)
}