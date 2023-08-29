package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntranceEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4 2.25a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0zM9.27 4H7.88a.73.73 0 0 0-.52.21l-4 4a.73.73 0 0 1-.51.21H1.73a.73.73 0 1 0 0 1.46h1.89a.73.73 0 0 0 .51-.21l4-4a.73.73 0 0 1 .48-.17h.66a.73.73 0 0 0 .73-.73a.73.73 0 0 0-.73-.77zm-4.52-.5a.75.75 0 0 0-.75.75V6l1.5-1.5v-.25a.75.75 0 0 0-.75-.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}