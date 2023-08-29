package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="15" r="3" fill="currentColor"/><path d="M18 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm-6 2a2 2 0 1 1-2 2a2 2 0 0 1 2-2zm0 16a5 5 0 1 1 5-5a5 5 0 0 1-5 5z" fill="currentColor"/>`),
		g.Group(children),
	)
}