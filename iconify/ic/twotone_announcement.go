package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneAnnouncement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4v13.17l.59-.59l.58-.58H20V4H4zm9 11h-2v-2h2v2zm0-4h-2V5h2v6z" opacity=".3"/><path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 14H5.17l-.59.59l-.58.58V4h16v12zM11 5h2v6h-2zm0 8h2v2h-2z"/>`),
		g.Group(children),
	)
}