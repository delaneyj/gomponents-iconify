package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M41 17H7a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V19a2 2 0 0 0-2-2ZM34 7H14v10h20V7Z"/><path stroke-linecap="round" d="M19 29h10m-5-5v10"/></g>`),
		g.Group(children),
	)
}