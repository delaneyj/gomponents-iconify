package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crosshairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 0h1v4L8 6l-.5-2V0zm1 16h-1v-4l.5-2l.5 2v4zM16 7.5v1h-4L10 8l2-.5h4zm-16 1v-1h4L6 8l-2 .5H0z"/><path fill="currentColor" d="M8 2.5a5.5 5.5 0 1 1 0 11A5.5 5.5 0 0 1 2.5 8a5.51 5.51 0 0 1 5.499-5.5zM8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1z"/>`),
		g.Group(children),
	)
}