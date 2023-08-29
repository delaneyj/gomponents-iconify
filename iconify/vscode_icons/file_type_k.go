package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#f2994a" d="m14.035 21.961l-1.884 1.884v5.977H6.917V2h5.234v15.415l1.014-1.3l5.017-5.887h6.285l-7.082 8.169l7.7 11.429H19.07Z"/>`),
		g.Group(children),
	)
}