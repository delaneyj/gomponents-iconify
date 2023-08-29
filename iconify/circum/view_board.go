package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.437 20.936H5.563a2.5 2.5 0 0 1-2.5-2.5V5.562a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5ZM5.563 4.062a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.562a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12.5 14.544a.5.5 0 0 1-1 0v-8a.5.5 0 0 1 1 0Zm4.217-2.091a.5.5 0 0 1-1 0V6.544a.5.5 0 0 1 1 0ZM8.28 6.544a.5.5 0 0 0-1 0v4a.5.5 0 0 0 1 0Z"/>`),
		g.Group(children),
	)
}