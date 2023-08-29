package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9a9 9 0 0 0-9-9Zm-7.5 9a7.44 7.44 0 0 1 1.7-4.74L16.74 17.8A7.49 7.49 0 0 1 4.5 12Zm13.3 4.74L7.26 6.2A7.49 7.49 0 0 1 17.8 16.74Z"/>`),
		g.Group(children),
	)
}