package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUserLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4H5v16h14V8h-4V4ZM3 2.992C3 2.444 3.447 2 3.998 2H16l5 5v13.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992Zm9 8.508a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM7.527 17a4.5 4.5 0 0 1 8.945 0H7.527Z"/>`),
		g.Group(children),
	)
}