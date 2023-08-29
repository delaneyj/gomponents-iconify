package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableInsertRowTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 23.5a.75.75 0 0 0 0 1.5h20.5a.75.75 0 0 0 0-1.5H3.75Zm0-20.5a.75.75 0 0 0 0 1.5h20.5a.75.75 0 0 0 0-1.5H3.75ZM3 16.25A2.75 2.75 0 0 0 5.75 19h16.5A2.75 2.75 0 0 0 25 16.25v-4.5A2.75 2.75 0 0 0 22.25 9H5.75A2.75 2.75 0 0 0 3 11.75v4.5Zm2.75 1.25c-.69 0-1.25-.56-1.25-1.25v-4.5c0-.69.56-1.25 1.25-1.25H10v7H5.75Zm5.75 0v-7h5v7h-5Zm6.5 0v-7h4.25c.69 0 1.25.56 1.25 1.25v4.5c0 .69-.56 1.25-1.25 1.25H18Z"/>`),
		g.Group(children),
	)
}