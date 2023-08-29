package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CassetteTapeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52H32a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h192a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12ZM72 196l16.8-22.4A4 4 0 0 1 92 172h72a4 4 0 0 1 3.2 1.6L184 196Zm156-4a4 4 0 0 1-4 4h-30l-20.4-27.2a12 12 0 0 0-9.6-4.8H92a12 12 0 0 0-9.6 4.8L62 196H32a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h192a4 4 0 0 1 4 4ZM176 84H80a28 28 0 0 0 0 56h96a28 28 0 0 0 0-56ZM60 112a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm39.57 20a27.94 27.94 0 0 0 0-40h56.86a27.94 27.94 0 0 0 0 40Zm76.43 0a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}