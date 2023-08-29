package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageTableTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3h.25v3.5H3v-.25A3.25 3.25 0 0 1 6.25 3ZM3 16V8h3.5v8H3Zm13-9.5H8V3h8v3.5Zm1.5 0V3h.25A3.25 3.25 0 0 1 21 6.25v.25h-3.5Zm0 1.5H21v8h-3.5V8Zm0 9.5H21v.25A3.25 3.25 0 0 1 17.75 21h-.25v-3.5Zm-1.5 0V21H8v-3.5h8Zm-9.5 0V21h-.25A3.25 3.25 0 0 1 3 17.75v-.25h3.5ZM8 15.068l2.409-2.409a2.25 2.25 0 0 1 3.182 0L16 15.068V8H8v7.068ZM14 9a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm.81 7l-2.28-2.28a.75.75 0 0 0-1.06 0L9.19 16h5.62Z"/>`),
		g.Group(children),
	)
}