package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ceilinglight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 64H576v512h128l320 448H0l320-448h128V64H288q-13 0-22.5-9.5T256 32t9.5-22.5T288 0h448q13 0 22.5 9.5T768 32t-9.5 22.5T736 64z"/>`),
		g.Group(children),
	)
}