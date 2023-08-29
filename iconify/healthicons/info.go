package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20ZM22 35a2 2 0 1 0 4 0V21a2 2 0 1 0-4 0v14Zm2-20a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}