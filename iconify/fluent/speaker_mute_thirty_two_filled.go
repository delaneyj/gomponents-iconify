package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerMuteThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 5.604c0-1.114-1.346-1.672-2.134-.884l-4.694 4.694A2 2 0 0 1 9.757 10H6a4 4 0 0 0-4 4v4a4 4 0 0 0 4 4h3.757a2 2 0 0 1 1.415.586l4.694 4.694c.788.788 2.134.23 2.134-.884V5.604Zm3.293 6.689a1 1 0 0 1 1.414 0L25 14.586l2.293-2.293a1 1 0 0 1 1.414 1.414L26.414 16l2.293 2.293a1 1 0 0 1-1.414 1.414L25 17.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L23.586 16l-2.293-2.293a1 1 0 0 1 0-1.414Z"/>`),
		g.Group(children),
	)
}