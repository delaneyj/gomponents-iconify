package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7.25h-.25V5A1.76 1.76 0 0 0 17 3.25a.67.67 0 0 0-.24 0l-11.9 4h-.27l-.17.06h-.14l-.16.09l-.12.17l-.14.12l-.11.1l-.12.15a.39.39 0 0 0-.08.1a1.62 1.62 0 0 0-.1.18l-.06.11a1.87 1.87 0 0 0-.07.22a.45.45 0 0 1 0 .11a1.87 1.87 0 0 0 0 .34v10A1.76 1.76 0 0 0 5 20.75h14A1.76 1.76 0 0 0 20.75 19V9A1.76 1.76 0 0 0 19 7.25Zm-1.92-2.49a.26.26 0 0 1 .17.24v2.25H9.62ZM19.25 19a.25.25 0 0 1-.25.25H5a.25.25 0 0 1-.25-.25V9A.25.25 0 0 1 5 8.75h14a.25.25 0 0 1 .25.25Z"/><circle cx="16.5" cy="14" r="1.25" fill="currentColor"/>`),
		g.Group(children),
	)
}