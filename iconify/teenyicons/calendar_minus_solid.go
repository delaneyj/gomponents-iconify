package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMinusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2ZM5 9h5V8H5v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}