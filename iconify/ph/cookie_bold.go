package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M167.31 160.69a16 16 0 1 1-22.62 0a16 16 0 0 1 22.62 0Zm-86.62-8a16 16 0 1 0 22.62 0a16 16 0 0 0-22.62 0Zm14.62-33.38a16 16 0 1 0-22.62 0a16 16 0 0 0 22.62 0Zm48-6.62a16 16 0 1 0 0 22.62a16 16 0 0 0 0-22.62ZM236 128A108 108 0 1 1 128 20a12 12 0 0 1 12 12a36 36 0 0 0 36 36a12 12 0 0 1 12 12a36 36 0 0 0 36 36a12 12 0 0 1 12 12Zm-24.67 10.65A60.17 60.17 0 0 1 165 91a60.17 60.17 0 0 1-47.66-46.32a84 84 0 1 0 94 94Z"/>`),
		g.Group(children),
	)
}