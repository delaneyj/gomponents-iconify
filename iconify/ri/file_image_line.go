package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 8V4H5v16h14V8h-4ZM3 2.992C3 2.444 3.447 2 3.998 2H16l5 5v13.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992ZM11 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6.5 7.5l-4-7L8 17h9.5Z"/>`),
		g.Group(children),
	)
}