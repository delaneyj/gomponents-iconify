package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 4.609A3.609 3.609 0 0 1 7.609 1H12a6 6 0 0 1 0 12H4V4.609ZM7.609 3C6.72 3 6 3.72 6 4.609V11h6a4 4 0 0 0 0-8H7.609Z" clip-rule="evenodd"/><path d="M4 13v6.941C4 21.631 5.37 23 7.059 23H14a6 6 0 0 0 2.102-11.621A5.979 5.979 0 0 1 12 12.998h2a4 4 0 0 1 0 8H7.059A1.059 1.059 0 0 1 6 19.942V13H4Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}