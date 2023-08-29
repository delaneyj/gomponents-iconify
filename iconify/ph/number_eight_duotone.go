package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 40v176a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M155.55 119.27a48 48 0 1 0-55.1 0a56 56 0 1 0 55.1 0ZM96 80a32 32 0 1 1 32 32a32 32 0 0 1-32-32Zm32 128a40 40 0 1 1 40-40a40 40 0 0 1-40 40Z"/></g>`),
		g.Group(children),
	)
}