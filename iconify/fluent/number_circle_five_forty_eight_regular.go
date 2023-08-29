package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M28.75 16.5h-6.832l-.332 4.978l.66-.005c1.028-.004 2.325.006 2.943.093a6.75 6.75 0 1 1-6.594 10.37a1.25 1.25 0 1 1 2.094-1.366a4.25 4.25 0 1 0 4.153-6.528c-.4-.057-1.485-.073-2.586-.069a114.314 114.314 0 0 0-1.818.023l-.12.003h-.04a1.251 1.251 0 0 1-1.277-1.332l.5-7.5A1.25 1.25 0 0 1 20.75 14h8.002a1.25 1.25 0 0 1 0 2.5ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4ZM6.5 24c0-9.665 7.835-17.5 17.5-17.5S41.5 14.335 41.5 24S33.665 41.5 24 41.5S6.5 33.665 6.5 24Z"/>`),
		g.Group(children),
	)
}