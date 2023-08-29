package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 68h-84l-26.67-20a20.12 20.12 0 0 0-12-4H40a20 20 0 0 0-20 20v136a20 20 0 0 0 20 20h176.89A19.13 19.13 0 0 0 236 200.89V88a20 20 0 0 0-20-20Zm-4 128H44V68h48l26.67 20a20.12 20.12 0 0 0 12 4H212Z"/>`),
		g.Group(children),
	)
}