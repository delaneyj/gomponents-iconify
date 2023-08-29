package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithoutMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FFCC4D" d="M36 18c0 9.941-8.059 18-18 18c-9.94 0-18-8.059-18-18C0 8.06 8.06 0 18 0c9.941 0 18 8.06 18 18"/><ellipse cx="11.5" cy="16.5" fill="#664500" rx="2.5" ry="3.5"/><ellipse cx="24.5" cy="16.5" fill="#664500" rx="2.5" ry="3.5"/>`),
		g.Group(children),
	)
}