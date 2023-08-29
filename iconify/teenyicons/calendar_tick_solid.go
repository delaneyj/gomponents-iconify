package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTickSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2Zm-6.476 9.232l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708l2.378 2.378Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}