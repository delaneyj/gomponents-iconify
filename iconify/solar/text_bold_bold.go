package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.609 1A3.609 3.609 0 0 0 4 4.609V19.94A3.06 3.06 0 0 0 7.059 23H14a6 6 0 0 0 2.102-11.621A6 6 0 0 0 12 1H7.609ZM12 11a4 4 0 0 0 0-8H7.609C6.72 3 6 3.72 6 4.609V11h6Zm-6 2v6.941C6 20.526 6.474 21 7.059 21H14a4 4 0 0 0 0-8H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}