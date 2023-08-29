package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luminous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 16V26"/><path d="M38.1421 21.8579L31.1421 28.8579"/><path d="M44 36H34"/><path d="M4 36H14"/><path d="M9.85791 21.8579L16.8579 28.8579"/><path d="M22 36H26"/></g>`),
		g.Group(children),
	)
}