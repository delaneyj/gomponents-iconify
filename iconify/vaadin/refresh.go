package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.6 5.6C3.5 3.5 5.6 2 8 2c3 0 5.4 2.2 5.9 5h2c-.5-3.9-3.8-7-7.9-7c-3 0-5.6 1.6-6.9 4.1L0 3v4h4L2.6 5.6zM16 9h-4.1l1.5 1.4c-.9 2.1-3 3.6-5.5 3.6C5 14 2.5 11.8 2 9H0c.5 3.9 3.9 7 7.9 7c3 0 5.6-1.7 7-4.1L16 13V9z"/>`),
		g.Group(children),
	)
}