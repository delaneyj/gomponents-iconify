package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.121 5.879A3 3 0 0 0 9 8a1 1 0 0 1-2 0a5 5 0 1 1 8.535 3.535L9.88 17.193A3 3 0 0 0 9.016 19H16a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1v-.686a5 5 0 0 1 1.464-3.536l5.657-5.657a3 3 0 0 0 0-4.242Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}