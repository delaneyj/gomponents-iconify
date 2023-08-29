package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningDiamondThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124 136V80a4 4 0 0 1 8 0v56a4 4 0 0 1-8 0Zm4 28a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm108-36a11.87 11.87 0 0 1-3.5 8.45l-96.05 96.06a12 12 0 0 1-16.9 0l-96-96.06a12 12 0 0 1 0-16.9l96.05-96.06a12 12 0 0 1 16.9 0l96.05 96.06A11.87 11.87 0 0 1 236 128Zm-8 0a3.9 3.9 0 0 0-1.16-2.79l-96.05-96.06a4 4 0 0 0-5.58 0l-96 96.06a3.94 3.94 0 0 0 0 5.58l96.05 96.06a4 4 0 0 0 5.58 0l96.05-96.06A3.9 3.9 0 0 0 228 128Z"/>`),
		g.Group(children),
	)
}