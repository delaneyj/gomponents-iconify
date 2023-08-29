package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FortAwesomeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 3C8.607 3 3 8.607 3 15.5S8.607 28 15.5 28S28 22.393 28 15.5S22.393 3 15.5 3zm0 2C21.29 5 26 9.71 26 15.5S21.29 26 15.5 26S5 21.29 5 15.5S9.71 5 15.5 5zM15 8v5h-1v-1h-1v1h-1v-1h-1v6H9v-1H8v2.49a8.523 8.523 0 0 0 6 4.37V20.5c0-.825.675-1.5 1.5-1.5s1.5.675 1.5 1.5v3.36a8.517 8.517 0 0 0 6-4.37V17h-1v1h-2v-6h-1v1h-1v-1h-1v1h-1v-3s.39-.23.9-.23c.5 0 .78.23 1.19.23c.55 0 .91-.23.91-.23V8s-.36.23-.91.23c-.41 0-.69-.23-1.19-.23c-.51 0-.9.23-.9.23V8h-1zm-2 7h1v2h-1v-2zm4 0h1v2h-1v-2z"/>`),
		g.Group(children),
	)
}