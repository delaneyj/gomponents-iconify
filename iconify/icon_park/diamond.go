package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10.6364 5H37.3636L45 18.3L24 43L3 18.3L10.6364 5Z" clip-rule="evenodd"/><path d="M10.6362 5L23.9999 43L37.3635 5"/><path d="M3 18.3H45"/><path d="M15.4092 18.3L24.0001 5L32.591 18.3"/></g>`),
		g.Group(children),
	)
}