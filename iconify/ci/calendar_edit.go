package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 10v10h14V10H5Zm0-4v2h14V6H5Zm4.8 13H8v-1.8l4.2-4.19l1.8 1.8L9.8 19Zm4.825-4.818l-1.8-1.8l1.375-1.369l1.8 1.8l-1.37 1.37l-.005-.001Z"/>`),
		g.Group(children),
	)
}