package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AidKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 4h-3V2c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v2H2C.9 4 0 4.9 0 6v8c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zM6 2h4v2H6V2zm6 9H9v3H7v-3H4V9h3V6h2v3h3v2z"/>`),
		g.Group(children),
	)
}