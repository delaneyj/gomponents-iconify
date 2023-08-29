package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 115.54V216h-64v-56a8 8 0 0 0-8-8h-32a8 8 0 0 0-8 8v56H40V115.54a8 8 0 0 1 2.62-5.92l80-75.54a8 8 0 0 1 10.77 0l80 75.54a8 8 0 0 1 2.61 5.92Z" opacity=".2"/><path d="M240 208h-16v-92.45a16 16 0 0 0-5.17-11.78l-80-75.48a1.14 1.14 0 0 1-.11-.11a16 16 0 0 0-21.53 0l-.11.11l-79.91 75.48A16 16 0 0 0 32 115.55V208H16a8 8 0 0 0 0 16h224a8 8 0 0 0 0-16ZM48 115.55l.11-.1L128 40l79.9 75.43l.11.1V208H160v-48a16 16 0 0 0-16-16h-32a16 16 0 0 0-16 16v48H48ZM144 208h-32v-48h32Z"/></g>`),
		g.Group(children),
	)
}