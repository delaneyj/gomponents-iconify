package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 15h32l-3 29H11L8 15Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M20.002 25.002v10m8-10.002v9.997"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15L28.324 3L36 15"/></g>`),
		g.Group(children),
	)
}