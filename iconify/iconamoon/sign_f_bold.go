package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignFBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 5v-.295C17 3.763 16.237 3 15.295 3v0a1.73 1.73 0 0 0-1.66 1.242L9.557 18.105A2.64 2.64 0 0 1 7.025 20H7m3.5-11H14"/>`),
		g.Group(children),
	)
}