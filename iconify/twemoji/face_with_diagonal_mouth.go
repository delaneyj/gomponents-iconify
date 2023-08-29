package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithDiagonalMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FFCC4D" d="M36 18c0 9.941-8.059 18-18 18S0 27.941 0 18S8.059 0 18 0s18 8.059 18 18"/><ellipse cx="11.5" cy="15.5" fill="#65471B" rx="2.5" ry="3.5"/><ellipse cx="24.5" cy="15.5" fill="#65471B" rx="2.5" ry="3.5"/><path fill="#65471B" d="m11.209 27.978l14-3a1.001 1.001 0 0 0-.419-1.957l-14 3a1.001 1.001 0 0 0 .419 1.957z"/>`),
		g.Group(children),
	)
}