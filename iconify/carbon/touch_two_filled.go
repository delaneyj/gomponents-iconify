package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 12a3.003 3.003 0 0 0-3 3v3h6v-3a3.003 3.003 0 0 0-3-3Z"/><path fill="currentColor" d="M16 6a9.01 9.01 0 0 0-9 9v13h18V15a9.01 9.01 0 0 0-9-9Zm5 14H11v-5a5 5 0 0 1 10 0Z"/><path fill="currentColor" d="M29 15h-2a11 11 0 0 0-22 0H3a13 13 0 0 1 26 0Z"/>`),
		g.Group(children),
	)
}