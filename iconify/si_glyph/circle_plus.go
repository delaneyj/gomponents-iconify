package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.021.097c-4.396 0-7.958 3.558-7.958 7.943c0 4.388 3.562 7.945 7.958 7.945c4.395 0 7.958-3.558 7.958-7.945c0-4.386-3.564-7.943-7.958-7.943zm2.304 8.985h-1.237v1.237c0 .979.059 1.769-1.088 1.769c-1.144 0-1.088-.79-1.088-1.769V9.082H6.675c-.979 0-1.769.056-1.769-1.088c0-1.146.79-1.088 1.769-1.088h1.237V5.669C7.912 4.69 7.856 3.9 9 3.9c1.146 0 1.088.79 1.088 1.769v1.237h1.237c.979 0 1.769-.059 1.769 1.088c0 1.144-.79 1.088-1.769 1.088z"/>`),
		g.Group(children),
	)
}