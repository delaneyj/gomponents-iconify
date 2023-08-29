package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-4H2v-6q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.675-.35-1.413-.55t-1.512-.2q-1.95 0-3.538 1.1T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H6ZM6 7V3h12v4H6Zm11.95 13.175L15.1 17.35l1.425-1.4l1.425 1.4l3.525-3.525l1.425 1.4l-4.95 4.95Z"/>`),
		g.Group(children),
	)
}