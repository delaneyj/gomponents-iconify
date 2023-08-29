package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FFCC4D" d="M36 18c0 9.941-8.059 18-18 18S0 27.941 0 18S8.059 0 18 0s18 8.059 18 18"/><ellipse cx="18" cy="25" fill="#664500" rx="4" ry="5"/><ellipse cx="12" cy="13.5" fill="#664500" rx="2.5" ry="3.5"/><ellipse cx="24" cy="13.5" fill="#664500" rx="2.5" ry="3.5"/>`),
		g.Group(children),
	)
}