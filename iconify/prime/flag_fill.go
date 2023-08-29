package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.42 4.45a.74.74 0 0 0-.69-.08a13.18 13.18 0 0 1-3.73 1a8.46 8.46 0 0 1-2.3-1a8.76 8.76 0 0 0-3-1.16c-1.29-.12-4.36.89-5 1.1a.75.75 0 0 0-.51.71V20a.75.75 0 0 0 1.5 0v-5.86a16 16 0 0 1 3.86-.85a8.47 8.47 0 0 1 2.4 1a9.11 9.11 0 0 0 2.82 1.13H15a16.37 16.37 0 0 0 4.21-1.13a.76.76 0 0 0 .48-.7V5.07a.74.74 0 0 0-.27-.62Z"/>`),
		g.Group(children),
	)
}