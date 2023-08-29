package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.05 3.636a1 1 0 0 1 0 1.414a7 7 0 0 0 0 9.9a1 1 0 1 1-1.414 1.414a9 9 0 0 1 0-12.728a1 1 0 0 1 1.414 0Zm9.9 0a1 1 0 0 1 1.414 0a9 9 0 0 1 0 12.728a1 1 0 1 1-1.414-1.414a7 7 0 0 0 0-9.9a1 1 0 0 1 0-1.414ZM7.879 6.464a1 1 0 0 1 0 1.414a3 3 0 0 0 0 4.243a1 1 0 1 1-1.415 1.414a5 5 0 0 1 0-7.07a1 1 0 0 1 1.415 0Zm4.242 0a1 1 0 0 1 1.415 0a5 5 0 0 1 0 7.072a1 1 0 0 1-1.415-1.415a3 3 0 0 0 0-4.242a1 1 0 0 1 0-1.415ZM10 9a1 1 0 0 1 1 1v.01a1 1 0 1 1-2 0V10a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}