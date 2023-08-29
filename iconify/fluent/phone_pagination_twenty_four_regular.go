package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonePaginationTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.75 2A2.25 2.25 0 0 1 18 4.25v15.5A2.25 2.25 0 0 1 15.75 22h-7.5A2.25 2.25 0 0 1 6 19.75V4.25A2.25 2.25 0 0 1 8.25 2h7.5Zm0 1.5h-7.5a.75.75 0 0 0-.75.75v15.5c0 .414.336.75.75.75h7.5a.75.75 0 0 0 .75-.75V4.25a.75.75 0 0 0-.75-.75ZM9.499 17.755a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm2.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm2.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}