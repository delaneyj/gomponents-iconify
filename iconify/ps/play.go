package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M18 427q4 2 9 2q11 0 15-6l170-192q12-16 0-28L42 11Q34-2 18 5Q5 9 5 24v384q0 15 13 19zM48 79l122 137L48 353V79z"/>`),
		g.Group(children),
	)
}