package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonThreeTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.58 1.5a1 1 0 0 0-.865.498l-.58 1a1 1 0 0 0 0 1.004l.58 1a1 1 0 0 0 .865.498h1.164a1 1 0 0 0 .865-.498l.58-1a1 1 0 0 0 0-1.004l-.58-1a1 1 0 0 0-.865-.498H3.58Zm0 1h1.164l.58 1l-.58 1H3.58L3 3.5l.58-1Zm3.465 1.998A1 1 0 0 1 7.91 4h1.164a1 1 0 0 1 .865.498l.58 1a1 1 0 0 1 0 1.004l-.58 1A1 1 0 0 1 9.074 8H7.91a1 1 0 0 1-.865-.498l-.58-1a1 1 0 0 1 0-1.004l.58-1ZM9.074 5H7.91l-.58 1l.58 1h1.164l.58-1l-.58-1ZM2.715 6.998A1 1 0 0 1 3.58 6.5h1.164a1 1 0 0 1 .865.498l.58 1a1 1 0 0 1 0 1.004l-.58 1a1 1 0 0 1-.865.498H3.58a1 1 0 0 1-.865-.498l-.58-1a1 1 0 0 1 0-1.004l.58-1Zm2.029.502H3.58L3 8.5l.58 1h1.164l.58-1l-.58-1Z"/>`),
		g.Group(children),
	)
}