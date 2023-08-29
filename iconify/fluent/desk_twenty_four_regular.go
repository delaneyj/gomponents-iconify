package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 4A2.25 2.25 0 0 0 2 6.25v11A2.75 2.75 0 0 0 4.75 20h4.5A2.75 2.75 0 0 0 12 17.25V9.5h8.5v9.75a.75.75 0 0 0 1.5 0v-13A2.25 2.25 0 0 0 19.75 4H4.25ZM3.5 9.5h7v7.75c0 .69-.56 1.25-1.25 1.25h-4.5c-.69 0-1.25-.56-1.25-1.25V9.5Zm0-1.5V6.25a.75.75 0 0 1 .75-.75h15.5a.75.75 0 0 1 .75.75V8h-17Zm2.25 3.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Z"/>`),
		g.Group(children),
	)
}