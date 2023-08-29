package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 68h-82.61l-26-29.29a20 20 0 0 0-15-6.71H40a20 20 0 0 0-20 20v148.62A19.41 19.41 0 0 0 39.38 220h177.51A19.13 19.13 0 0 0 236 200.89V88a20 20 0 0 0-20-20ZM44 56h46.61l10.67 12H44Zm168 140H44V92h168Z"/>`),
		g.Group(children),
	)
}