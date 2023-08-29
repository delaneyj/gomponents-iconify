package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.983 1.25H9a.75.75 0 0 0 0 1.5h4.992l-5.55 18.5H3a.75.75 0 1 0 0 1.5h5.983a.98.98 0 0 0 .034 0H15a.75.75 0 0 0 0-1.5h-4.992l5.55-18.5H21a.75.75 0 0 0 0-1.5h-6.017Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}