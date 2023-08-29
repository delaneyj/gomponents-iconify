package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderUserBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220.51 197.94a36 36 0 1 0-57 0a43.75 43.75 0 0 0-15.08 23a12 12 0 0 0 8.52 14.67a11.77 11.77 0 0 0 3.05.39a12 12 0 0 0 11.59-8.92C174 218.2 182.35 212 192 212s18 6.2 20.4 15.08a12 12 0 0 0 23.19-6.17a43.7 43.7 0 0 0-15.08-22.97ZM192 164a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm24-96h-82.61l-26-29.29A20 20 0 0 0 92.41 32H40a20 20 0 0 0-20 20v148.61A19.41 19.41 0 0 0 39.38 220h73.18a12 12 0 0 0 0-24H44V92h168v16a12 12 0 0 0 24 0V88a20 20 0 0 0-20-20ZM44 68V56h46.61l10.67 12Z"/>`),
		g.Group(children),
	)
}