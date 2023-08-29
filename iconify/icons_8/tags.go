package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m14.594 4l-.313.28l-11 11l-.685.72l.687.72l9 9l.72.686l.72-.687l11-11l.28-.315V4H14.594zm.844 2H23v7.563l-10 10L5.437 16l10-10zM26 7v2h1v8.156l-9.5 9.438l-1.25-1.25l-1.406 1.406l1.937 1.97l.72.686l.69-.687L28.72 18.31L29 18V7h-3zm-6 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}