package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Needle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21c-.667-.667 3.262-6.236 11.785-16.709a3.5 3.5 0 1 1 5.078 4.791C9.288 17.694 3.667 21.667 3 21zM17.5 6.5l-1 1"/>`),
		g.Group(children),
	)
}