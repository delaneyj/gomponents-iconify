package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderLockBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 152h-4v-4a32 32 0 0 0-64 0v4h-4a12 12 0 0 0-12 12v44a12 12 0 0 0 12 12h72a12 12 0 0 0 12-12v-44a12 12 0 0 0-12-12Zm-44-4a8 8 0 0 1 16 0v4h-16Zm32 48h-48v-20h48Zm4-128h-82.61l-26-29.29a20 20 0 0 0-15-6.71H40a20 20 0 0 0-20 20v148.62A19.41 19.41 0 0 0 39.38 220h65.18a12 12 0 0 0 0-24H44V92h168a12 12 0 0 0 24 0v-4a20 20 0 0 0-20-20ZM44 68V56h46.61l10.67 12Z"/>`),
		g.Group(children),
	)
}