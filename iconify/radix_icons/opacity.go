package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opacity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1.5C4.5 4.25 3 6.5 3 9a4.5 4.5 0 1 0 9 0c0-2.5-1.5-4.75-4.5-7.5ZM11 9c0-1.888-1.027-3.728-3.5-6.126C5.027 5.272 4 7.112 4 9a3.5 3.5 0 1 0 7 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}