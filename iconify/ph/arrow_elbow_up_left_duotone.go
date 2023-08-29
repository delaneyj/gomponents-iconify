package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowUpLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M96 32v96L48 80Z" opacity=".2"/><path d="M192 72h-88V32a8 8 0 0 0-13.66-5.66l-48 48a8 8 0 0 0 0 11.32l48 48A8 8 0 0 0 104 128V88h80v136a8 8 0 0 0 16 0V80a8 8 0 0 0-8-8ZM88 108.69L59.31 80L88 51.31Z"/></g>`),
		g.Group(children),
	)
}