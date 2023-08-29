package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningDiamondBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 68a12 12 0 0 1 12 12v52a12 12 0 0 1-24 0V80a12 12 0 0 1 12-12Zm0 88a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm116-28a19.86 19.86 0 0 1-5.84 14.11l-96 96.06a20 20 0 0 1-28.21 0l-96-96.06a20 20 0 0 1 0-28.22L114 17.83a20 20 0 0 1 28.21 0l96.06 96.06A19.86 19.86 0 0 1 244 128Zm-25.68 0L128 37.67L37.68 128L128 218.33Z"/>`),
		g.Group(children),
	)
}