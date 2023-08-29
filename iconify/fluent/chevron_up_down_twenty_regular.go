package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.911 7.216a.5.5 0 0 1-.765.638L10 3.707L5.854 7.854l-.07.057a.5.5 0 0 1-.638-.765l4.5-4.5l.07-.057a.5.5 0 0 1 .638.057l4.5 4.5l.057.07ZM5.09 12.784a.5.5 0 0 1 .765-.638L10 16.293l4.146-4.147l.07-.057a.5.5 0 0 1 .638.765l-4.5 4.5l-.07.057a.5.5 0 0 1-.638-.057l-4.5-4.5l-.057-.07Z"/>`),
		g.Group(children),
	)
}