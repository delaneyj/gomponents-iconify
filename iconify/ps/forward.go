package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M13 427q4 2 8 2q12 0 15-6l135-149v134q0 15 12 19q1 0 4 1t5 1q12 0 15-6l171-192q10-16 0-28L207 11q-7-14-24-6q-12 4-12 19v137L36 11Q29-3 13 5Q0 9 0 24v384q0 15 13 19zM213 79l122 137l-122 137V79zM43 79l121 137L43 353V79z"/>`),
		g.Group(children),
	)
}