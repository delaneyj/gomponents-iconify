package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H7a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 7 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.77.77 0 0 1-.42.68a.78.78 0 0 1-.33.07Zm-9.25-5.5H12a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.16H7.75Z"/>`),
		g.Group(children),
	)
}