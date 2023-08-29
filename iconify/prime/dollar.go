package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v16a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M13.5 18.75H7a.75.75 0 0 1 0-1.5h6.5A2.54 2.54 0 0 0 16.25 15a2.54 2.54 0 0 0-2.75-2.25h-3A4 4 0 0 1 6.25 9a4 4 0 0 1 4.25-3.75H16a.75.75 0 0 1 0 1.5h-5.5A2.54 2.54 0 0 0 7.75 9a2.54 2.54 0 0 0 2.75 2.25h3A4 4 0 0 1 17.75 15a4 4 0 0 1-4.25 3.75Z"/>`),
		g.Group(children),
	)
}