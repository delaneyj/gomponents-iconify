package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipAttechmentHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6h9.75a5.25 5.25 0 1 1 0 10.5H5.5a3.5 3.5 0 1 1 0-7h11.25a1.75 1.75 0 1 1 0 3.5H7"/>`),
		g.Group(children),
	)
}