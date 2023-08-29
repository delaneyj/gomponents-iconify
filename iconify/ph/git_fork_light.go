package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitForkLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 64a30 30 0 1 0-36 29.4V104a18 18 0 0 1-18 18H88a18 18 0 0 1-18-18V93.4a30 30 0 1 0-12 0V104a30 30 0 0 0 30 30h34v28.6a30 30 0 1 0 12 0V134h34a30 30 0 0 0 30-30V93.4A30.05 30.05 0 0 0 222 64ZM46 64a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm100 128a18 18 0 1 1-18-18a18 18 0 0 1 18 18Zm46-110a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}