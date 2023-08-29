package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 16a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1 0-1.5h3Zm3-5a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1 0-1.5h9Zm3-5a.75.75 0 0 1 0 1.5h-15a.75.75 0 0 1 0-1.5h15Z"/>`),
		g.Group(children),
	)
}