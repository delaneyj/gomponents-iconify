package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v15c0 4 3 6 6 6s6-2 6-6V6c0-3-2-4-4-4s-4 1-4 4v17c0 1 1 2 2 2s2-1 2-2V9"/>`),
		g.Group(children),
	)
}