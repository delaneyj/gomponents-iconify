package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeBasics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.762 36.581c-2.992-2.639-4.84-6.422-4.84-10.645c0-5.455 3.168-10.206 7.743-12.582m7.567-1.408a14.005 14.005 0 0 1 12.845 13.99c0 7.742-6.334 14.077-14.077 14.077c-1.32 0-2.64-.176-3.871-.528M11.331 44.5l17.508-41"/>`),
		g.Group(children),
	)
}