package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3 9.5h3m8 8a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm9.5-11h-3a3 3 0 0 1-3-3h-7a3 3 0 0 1-3 3h-7v14h20a3 3 0 0 0 3-3v-11Z"/>`),
		g.Group(children),
	)
}