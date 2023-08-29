package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitnessCentreEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7 6H4V5h3zM2 3v1H1v1a.5.5 0 0 0 0 1v1h1v1h1.5V3zm8 2V4H9V3H7.5v5H9V7h1V6a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}