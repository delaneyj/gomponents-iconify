package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerMuteThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 5.604c0-1.114-1.346-1.672-2.134-.884l-4.694 4.694A2 2 0 0 1 9.757 10H6a4 4 0 0 0-4 4v4a4 4 0 0 0 4 4h3.757a2 2 0 0 1 1.415.586l4.694 4.694c.788.788 2.134.23 2.134-.884V5.604Zm-5.414 5.224L16 7.415v17.172l-3.414-3.414A4 4 0 0 0 9.757 20H6a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h3.757a4 4 0 0 0 2.829-1.171Zm10.121 1.465a1 1 0 0 0-1.414 1.414L23.586 16l-2.293 2.293a1 1 0 0 0 1.414 1.414L25 17.414l2.293 2.293a1 1 0 0 0 1.414-1.414L26.414 16l2.293-2.293a1 1 0 0 0-1.414-1.414L25 14.586l-2.293-2.293Z"/>`),
		g.Group(children),
	)
}