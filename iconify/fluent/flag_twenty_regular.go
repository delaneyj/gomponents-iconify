package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 13h11a.5.5 0 0 0 .416-.777L13.101 8l2.815-4.223A.5.5 0 0 0 15.5 3H4a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 1 0V13Zm0-1V4h10.066l-2.482 3.723a.5.5 0 0 0 0 .554L14.566 12H4.5Z"/>`),
		g.Group(children),
	)
}