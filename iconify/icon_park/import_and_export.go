package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportAndExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="3.99"><path d="M14 25.9999L5 34.9999L14 43.9999"/><path d="M5 35.0083H22.5"/><path d="M34.0005 18L43.0005 27L34.0005 36"/><path d="M43 27.0084H25.5"/><path d="M4.5 24V7.5L43.5 7.5V15"/></g>`),
		g.Group(children),
	)
}