package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25 2.75a.75.75 0 0 0-1.5 0v22.5a.75.75 0 0 0 1.5 0V2.75ZM19.25 5A2.75 2.75 0 0 1 22 7.75v2.5A2.75 2.75 0 0 1 19.25 13H5.75A2.75 2.75 0 0 1 3 10.25v-2.5A2.75 2.75 0 0 1 5.75 5h13.5Zm1.25 2.75c0-.69-.56-1.25-1.25-1.25H5.75c-.69 0-1.25.56-1.25 1.25v2.5c0 .69.56 1.25 1.25 1.25h13.5c.69 0 1.25-.56 1.25-1.25v-2.5ZM19.25 15A2.75 2.75 0 0 1 22 17.75v2.5A2.75 2.75 0 0 1 19.25 23h-8a2.75 2.75 0 0 1-2.75-2.75v-2.5A2.75 2.75 0 0 1 11.25 15h8Zm1.25 2.75c0-.69-.56-1.25-1.25-1.25h-8c-.69 0-1.25.56-1.25 1.25v2.5c0 .69.56 1.25 1.25 1.25h8c.69 0 1.25-.56 1.25-1.25v-2.5Z"/>`),
		g.Group(children),
	)
}