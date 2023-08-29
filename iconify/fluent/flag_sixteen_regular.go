package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 9V3h8.028l-1.935 2.71a.5.5 0 0 0 0 .58L12.028 9H4Zm0 1h9a.5.5 0 0 0 .407-.79L11.114 6l2.293-3.21A.5.5 0 0 0 13 2H3.5a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 1 0V10Z"/>`),
		g.Group(children),
	)
}