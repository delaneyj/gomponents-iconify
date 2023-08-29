package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.75a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75Zm0-9.5a.76.76 0 0 1-.75-.75V7a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}