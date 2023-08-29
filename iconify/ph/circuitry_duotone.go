package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitryDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 40H48a8 8 0 0 0-8 8v160a8 8 0 0 0 8 8h160a8 8 0 0 0 8-8V48a8 8 0 0 0-8-8ZM88 184a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm80-64a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z" opacity=".2"/><path d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM88 160a8 8 0 1 1-8 8a8 8 0 0 1 8-8ZM48 48h32v97.38a24 24 0 1 0 16 0v-30.07l48 48V208H48Zm160 160h-48v-48a8 8 0 0 0-2.34-5.66L96 92.69V48h32v24a8 8 0 0 0 2.34 5.66l16 16A23.74 23.74 0 0 0 144 104a24 24 0 1 0 24-24a23.74 23.74 0 0 0-10.34 2.35L144 68.69V48h64v160ZM168 96a8 8 0 1 1-8 8a8 8 0 0 1 8-8Z"/></g>`),
		g.Group(children),
	)
}