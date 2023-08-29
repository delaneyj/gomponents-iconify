package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskBackBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H91.31a19.86 19.86 0 0 0-14.14 5.86L33.86 77.17A19.86 19.86 0 0 0 28 91.31V208a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V93l41-41h71v24h-64a12 12 0 0 0 0 24h68a20 20 0 0 0 20-20V52h16Zm-76-88a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 48a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}