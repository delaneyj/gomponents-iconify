package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DvLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.608 3H21a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-7v-2h6V5h-6.255A6.968 6.968 0 0 1 15 9a6.992 6.992 0 0 1-3 5.745V21a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-6.255A7 7 0 0 1 11.608 3ZM6 13.584V20h4v-6.416A5.001 5.001 0 0 0 8 4a5 5 0 0 0-2 9.584ZM8 12a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm9-3h2v2h-2V7ZM7 17h2v2H7v-2Z"/>`),
		g.Group(children),
	)
}