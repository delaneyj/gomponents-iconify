package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.4 20h1.325L20 18.425l-4.875-4.875L13.35 8.2l-3.325.75L9 7.975v-3.35l-.7-.35L4.4 9.5h3l11 10.5Zm-5.025 0H15.5l-8.9-8.5H4.05l9.325 8.5Zm-.75 2L.675 11.15L7.7 1.725l3.3 1.65v3.35l3.675-.95l2.2 6.675L22 17.575L21.25 22h-8.625Z"/>`),
		g.Group(children),
	)
}