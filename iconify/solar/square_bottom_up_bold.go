package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareBottomUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m5.568 15.25l-3.205 3.205C2 16.991 2 14.961 2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22c-2.96 0-4.991 0-6.455-.363l3.205-3.205V19a2.25 2.25 0 0 0 4.5 0v-6A2.25 2.25 0 0 0 11 10.75H5a2.25 2.25 0 0 0 0 4.5h.568Z"/><path d="M5 12.25a.75.75 0 0 0 0 1.5h4.19l-6.72 6.72a.75.75 0 1 0 1.06 1.06l6.72-6.72V19a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-.75-.75H5Z"/></g>`),
		g.Group(children),
	)
}