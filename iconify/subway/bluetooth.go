package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M375.4 157.5L217.8 0v196.9L139 118.1l-39.4 39.4l98.5 98.5l-98.5 98.5l39.4 39.4l78.8-78.8V512l157.5-157.5l-98.4-98.5l98.5-98.5zm-118.2-59l59.1 59.1l-59.1 59.1V98.5zm59.1 256l-59.1 59.1V295.4l59.1 59.1z"/>`),
		g.Group(children),
	)
}