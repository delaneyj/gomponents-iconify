package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M152 80H56l48-48Z" opacity=".2"/><path d="M200 216a88.1 88.1 0 0 1-88-88V88h40a8 8 0 0 0 5.66-13.66l-48-48a8 8 0 0 0-11.32 0l-48 48A8 8 0 0 0 56 88h40v40a104.11 104.11 0 0 0 104 104a8 8 0 0 0 0-16ZM104 43.31L132.69 72H75.31Z"/></g>`),
		g.Group(children),
	)
}