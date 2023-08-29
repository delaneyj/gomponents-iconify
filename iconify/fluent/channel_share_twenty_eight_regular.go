package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChannelShareTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v7.5A3.75 3.75 0 0 0 6.75 18h7.42a3.001 3.001 0 1 0-.129-1.5H6.75a2.25 2.25 0 0 1-2.25-2.25v-7.5A2.25 2.25 0 0 1 6.75 4.5h7.5a2.25 2.25 0 0 1 2.25 2.25v.5a.75.75 0 0 0 1.5 0v-.5A3.75 3.75 0 0 0 14.25 3h-7.5ZM17 15.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-3.041-4h7.291a2.25 2.25 0 0 1 2.25 2.25v7.5a2.25 2.25 0 0 1-2.25 2.25h-7.5a2.25 2.25 0 0 1-2.25-2.25v-.5a.75.75 0 0 0-1.5 0v.5A3.75 3.75 0 0 0 13.75 25h7.5A3.75 3.75 0 0 0 25 21.25v-7.5A3.75 3.75 0 0 0 21.25 10h-7.42a3.001 3.001 0 1 0 .129 1.5ZM9.5 11a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}