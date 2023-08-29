package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 1h8v2H4V1zm11 3H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h3v4h8v-4h3c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1zM2 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm9 7H5V9h6v5z"/>`),
		g.Group(children),
	)
}