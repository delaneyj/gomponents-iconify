package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 16h3L9 0H7L1 16h3l1.9-5h4.2l1.9 5zM6.7 9L8 5.4L9.3 9H6.7z"/>`),
		g.Group(children),
	)
}