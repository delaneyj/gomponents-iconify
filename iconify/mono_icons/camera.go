package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.293 4.293A1 1 0 0 1 9 4h6a1 1 0 0 1 .707.293L17.414 6H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2.586l1.707-1.707zM9.414 6L7.707 7.707A1 1 0 0 1 7 8H4v10h16V8h-3a1 1 0 0 1-.707-.293L14.586 6H9.414zM12 10.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0z"/>`),
		g.Group(children),
	)
}