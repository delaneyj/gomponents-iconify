package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.67 8.14a2 2 0 0 0-1.42-1.43A48.44 48.44 0 0 0 12 6.38a48.44 48.44 0 0 0-6.25.33a2 2 0 0 0-1.42 1.43A21.27 21.27 0 0 0 4 12a21.42 21.42 0 0 0 .33 3.88a2 2 0 0 0 1.42 1.4a48.44 48.44 0 0 0 6.25.33a48.44 48.44 0 0 0 6.25-.33a2 2 0 0 0 1.42-1.4A21.42 21.42 0 0 0 20 12a21.27 21.27 0 0 0-.33-3.86Zm-9.31 6.25V9.63L14.55 12l-4.19 2.38Z"/>`),
		g.Group(children),
	)
}