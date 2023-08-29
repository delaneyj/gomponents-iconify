package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M197 365q22 0 22-21V88q0-21-22-21q-21 0-21 21v73L42 11Q34-2 18 5Q5 9 5 24v384q0 15 13 19q4 2 9 2q11 0 15-6l134-149v70q0 21 21 21zM48 353V79l122 137z"/>`),
		g.Group(children),
	)
}