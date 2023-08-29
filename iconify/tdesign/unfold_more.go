package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.586L17.414 9L16 10.414l-4-4l-4 4L6.586 9L12 3.586Zm-4 10l4 4l4-4L17.414 15L12 20.414L6.586 15L8 13.586Z"/>`),
		g.Group(children),
	)
}