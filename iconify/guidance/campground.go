package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campground(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18 4.5h1.5c0 4.92 1.247 9.76 3.623 14.067l.377.683v.25H14m5 0V14m-.5-9.5h-14c0 4.92-1.246 9.76-3.623 14.067L.5 19.25v.25h13.85l.15-.25l.377-.683A29.12 29.12 0 0 0 18.5 4.5Z"/>`),
		g.Group(children),
	)
}