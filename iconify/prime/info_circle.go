package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.75a.76.76 0 0 1-.75-.75v-5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75Zm0-7.5a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9Zm0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5Z"/>`),
		g.Group(children),
	)
}