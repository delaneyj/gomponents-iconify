package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buddy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.932 7.078L17.26.338a2.508 2.508 0 0 0-2.521 0L3.067 7.072a2.527 2.527 0 0 0-1.26 2.187v13.479c0 .896.479 1.729 1.26 2.182l11.672 6.74c.781.448 1.74.448 2.521 0l11.677-6.74a2.53 2.53 0 0 0 1.26-2.182V9.259c0-.901-.484-1.734-1.26-2.182zM19 17.411l-5.073 5.073l-1.802-1.802l5.073-5.073l-5.073-5.073l1.802-1.802l6.875 6.875z"/>`),
		g.Group(children),
	)
}