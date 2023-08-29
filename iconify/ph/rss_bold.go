package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M101.74 154.26A75.53 75.53 0 0 1 124 208a12 12 0 0 1-24 0a52 52 0 0 0-52-52a12 12 0 0 1 0-24a75.51 75.51 0 0 1 53.74 22.26ZM48 84a12 12 0 0 0 0 24a100 100 0 0 1 100 100a12 12 0 0 0 24 0A124 124 0 0 0 48 84Zm121.62 2.38A170.85 170.85 0 0 0 48 36a12 12 0 0 0 0 24a147 147 0 0 1 104.65 43.35A147 147 0 0 1 196 208a12 12 0 0 0 24 0a170.85 170.85 0 0 0-50.38-121.62ZM52 188a16 16 0 1 0 16 16a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}