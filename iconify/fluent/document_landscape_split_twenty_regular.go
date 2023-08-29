package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLandscapeSplitTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 16a2 2 0 0 0 2-2V9.414a1.5 1.5 0 0 0-.44-1.06l-3.914-3.915A1.5 1.5 0 0 0 12.586 4H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12Zm1-2a1 1 0 0 1-1 1h-6V5h2v3.5a1.5 1.5 0 0 0 1.5 1.5H17v4ZM9 5v10H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h5Zm4 3.5V5.207L16.793 9H13.5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}