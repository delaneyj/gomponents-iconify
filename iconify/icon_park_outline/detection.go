package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Detection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="m6.45 34.85l3.99-3.348l.513-5.87l7.98-6.696l8.778-7.365l3.192-2.679a7.292 7.292 0 1 1 9.373 11.172l-3.191 2.678l-8.778 7.365l-7.98 6.696l-5.867-.51l-3.987 3.345a3.129 3.129 0 0 1-4.408-.386a3.125 3.125 0 0 1 .385-4.403Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M23.436 9.718L38.209 27.32M18.576 29l5.726-5"/></g>`),
		g.Group(children),
	)
}