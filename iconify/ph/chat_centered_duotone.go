package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCenteredDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 56v128a8 8 0 0 1-8 8h-59.47a8 8 0 0 0-6.86 3.88l-14.81 24.24a8 8 0 0 1-13.72 0l-14.81-24.24a8 8 0 0 0-6.86-3.88H40a8 8 0 0 1-8-8V56a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M216 40H40a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h59.47l14.81 24.23a16 16 0 0 0 27.41.06L156.53 200H216a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm0 144h-59.47a16.07 16.07 0 0 0-13.69 7.71L128 216l-14.85-24.3a16.08 16.08 0 0 0-13.68-7.7H40V56h176Z"/></g>`),
		g.Group(children),
	)
}