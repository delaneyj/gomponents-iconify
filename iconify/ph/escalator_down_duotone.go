package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 160v40a8 8 0 0 1-8 8h-52.5a8 8 0 0 1-5.88-2.57l-91.24-98.86A8 8 0 0 0 68.5 104H32a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h52.5a8 8 0 0 1 5.88 2.57l91.24 98.86a8 8 0 0 0 5.88 2.57H224a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M170.34 85.66a8 8 0 0 1 11.32-11.32L192 84.69V48a8 8 0 0 1 16 0v36.69l10.34-10.35a8 8 0 0 1 11.32 11.32l-24 24a8 8 0 0 1-11.32 0ZM240 160v40a16 16 0 0 1-16 16h-52.5a16.06 16.06 0 0 1-11.76-5.15L68.5 112H32a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h52.5a16.06 16.06 0 0 1 11.76 5.15L187.5 144H224a16 16 0 0 1 16 16Zm-16 0h-36.5a16.06 16.06 0 0 1-11.76-5.15L84.5 56H32v40h36.5a16.06 16.06 0 0 1 11.76 5.15L171.5 200H224Z"/></g>`),
		g.Group(children),
	)
}