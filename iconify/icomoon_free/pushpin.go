package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.5 0L7 1.5L8.5 3L5 7H1.5l2.75 2.75L0 15.385V16h.615l5.635-4.25L9 14.5V11l4-3.5L14.5 9L16 7.5L8.5 0zM7 8.5l-1-1L9.5 4l1 1L7 8.5z"/>`),
		g.Group(children),
	)
}