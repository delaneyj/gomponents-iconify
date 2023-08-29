package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 25C4 18.3502 9.39624 13 16 13H44"/><path d="M38 7L44 13L38 19"/><path d="M44 23C44 29.6498 38.6038 35 32 35H4"/><path d="M10 41L4 35L10 29"/></g>`),
		g.Group(children),
	)
}