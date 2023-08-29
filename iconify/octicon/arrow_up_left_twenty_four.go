package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.75 15.5a.75.75 0 0 1-.75-.75v-9A.75.75 0 0 1 5.75 5h9a.75.75 0 0 1 0 1.5H7.56l10.22 10.22a.749.749 0 0 1-.326 1.275a.749.749 0 0 1-.734-.215L6.5 7.56v7.19a.75.75 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}