package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMinusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 144a12 12 0 0 1 12-12h48a12 12 0 0 1 0 24h-48a12 12 0 0 1-12-12Zm144-56v112.89A19.13 19.13 0 0 1 216.89 220H39.38A19.41 19.41 0 0 1 20 200.62V52a20 20 0 0 1 20-20h52.41a20 20 0 0 1 15 6.71l26 29.29H216a20 20 0 0 1 20 20ZM44 68h57.28L90.61 56H44Zm168 24H44v104h168Z"/>`),
		g.Group(children),
	)
}