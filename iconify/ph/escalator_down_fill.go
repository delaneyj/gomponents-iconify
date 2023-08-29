package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M170.34 85.66a8 8 0 0 1 11.32-11.32L192 84.69V48a8 8 0 0 1 16 0v36.69l10.34-10.35a8 8 0 0 1 11.32 11.32l-24 24a8 8 0 0 1-11.32 0ZM224 144h-36.5L96.26 45.15A16.06 16.06 0 0 0 84.5 40H32a16 16 0 0 0-16 16v40a16 16 0 0 0 16 16h36.5l91.24 98.85A16.06 16.06 0 0 0 171.5 216H224a16 16 0 0 0 16-16v-40a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}