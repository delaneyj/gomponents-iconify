package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRectangleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 56v144a8 8 0 0 1-8 8h-20.1a72 72 0 0 0-67.9-48a40 40 0 1 0-40-40a40 40 0 0 0 40 40a72 72 0 0 0-67.9 48H40a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16ZM96 120a32 32 0 1 1 32 32a32 32 0 0 1-32-32Zm-23.43 80a64 64 0 0 1 110.86 0ZM216 200h-14.67a80.14 80.14 0 0 0-43.69-42.28a48 48 0 1 0-59.28 0A80.14 80.14 0 0 0 54.67 200H40V56h176v144Z"/></g>`),
		g.Group(children),
	)
}