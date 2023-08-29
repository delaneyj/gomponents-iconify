package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.67.81v1.48A.69.69 0 0 1 8 3H3.09a.69.69 0 0 1-.7-.7V.81a.345.345 0 0 1 .69 0v.69H4V1a.5.5 0 0 1 1 0v.5h1V1a.5.5 0 0 1 1 0v.5h1V.81a.34.34 0 1 1 .67 0zm1.39 8.82a.35.35 0 0 1-.35.35H1.35a.35.35 0 0 1 0-.7h.35a.68.68 0 0 0 .7-.7s.7-3.2.7-3.89A.68.68 0 0 1 3.79 4h3.48a.68.68 0 0 1 .7.7c0 .7.7 3.89.7 3.89a.68.68 0 0 0 .7.7h.34a.35.35 0 0 1 .35.35v-.01zM6.5 7.5a1 1 0 1 0-2 0v2h2v-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}