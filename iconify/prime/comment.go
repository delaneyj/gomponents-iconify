package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 20.25a.75.75 0 0 1-.72-1L5.38 14a7.76 7.76 0 0 1-.52-2.83a8 8 0 0 1 .62-3.1a8.12 8.12 0 0 1 1.7-2.52a7.83 7.83 0 0 1 2.53-1.7a7.92 7.92 0 0 1 6.19 0a8 8 0 0 1 4.85 7.32a8 8 0 0 1-2.33 5.62a8.12 8.12 0 0 1-2.52 1.7a8 8 0 0 1-5.93.1l-5.25 1.6a.83.83 0 0 1-.22.06Zm8.3-15.5a6.49 6.49 0 0 0-5.94 3.94a6.55 6.55 0 0 0 0 5a.75.75 0 0 1 0 .51l-1.23 4.17l4.15-1.26a.75.75 0 0 1 .51 0a6.52 6.52 0 0 0 5 0a6.44 6.44 0 0 0 3.43-8.45a6.45 6.45 0 0 0-5.92-3.91Z"/>`),
		g.Group(children),
	)
}