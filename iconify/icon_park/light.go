package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Light(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 16V22"/><path d="M38.1421 21.8579L33.8994 26.1005"/><path d="M44 36H38"/><path d="M4 36H10"/><path d="M9.85791 21.8579L14.1006 26.1005"/><path d="M18 36H30"/></g>`),
		g.Group(children),
	)
}