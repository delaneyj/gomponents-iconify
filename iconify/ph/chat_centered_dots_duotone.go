package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCenteredDotsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 56v128a8 8 0 0 1-8 8h-59.47a8 8 0 0 0-6.86 3.88l-14.81 24.24a8 8 0 0 1-13.72 0l-14.81-24.24a8 8 0 0 0-6.86-3.88H40a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M116 120a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm56 12a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm60-76v128a16 16 0 0 1-16 16h-59.47l-14.84 24.29a16 16 0 0 1-27.41-.06L99.47 200H40a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16Zm-16 0H40v128h59.47a16.08 16.08 0 0 1 13.7 7.73L128 216l14.82-24.32a16.07 16.07 0 0 1 13.71-7.68H216ZM84 132a12 12 0 1 0-12-12a12 12 0 0 0 12 12Z"/></g>`),
		g.Group(children),
	)
}