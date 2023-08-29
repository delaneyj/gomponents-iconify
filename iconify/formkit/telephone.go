package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telephone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.38 1.52c-.44.05-.79.41-.85.85l-.77 3.27c-.11.46.12.94.55 1.14l.59.27C9.46 8.2 8.49 8.99 8.49 8.99s-.8.97-1.94 1.41l-.27-.59a.998.998 0 0 0-1.14-.55l-3.27.77c-.44.06-.8.41-.85.85c-.1.82-.07 2.1.85 2.78c0 0 4.15 2.92 9.19-2.12c5.04-5.04 2.12-9.19 2.12-9.19c-.69-.92-1.97-.94-2.78-.85Z"/>`),
		g.Group(children),
	)
}