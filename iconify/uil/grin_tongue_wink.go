package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinTongueWink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm5.62-10.87a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41ZM9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm6 2H9a1 1 0 0 0 0 2a3 3 0 0 0 6 0a1 1 0 0 0 0-2Zm-3 3a1 1 0 0 1-1-1h2a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}