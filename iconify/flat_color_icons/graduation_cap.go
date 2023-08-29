package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#37474F"><path d="M9 20h30v13H9z"/><ellipse cx="24" cy="33" rx="15" ry="6"/></g><path fill="#78909C" d="M23.1 8.2L.6 18.1c-.8.4-.8 1.5 0 1.9l22.5 9.9c.6.2 1.2.2 1.8 0L47.4 20c.8-.4.8-1.5 0-1.9L24.9 8.2c-.6-.3-1.2-.3-1.8 0z"/><g fill="#37474F"><path d="m43.2 20.4l-20-3.4c-.5-.1-1.1.3-1.2.8c-.1.5.3 1.1.8 1.2L42 22.2V37c0 .6.4 1 1 1s1-.4 1-1V21.4c0-.5-.4-.9-.8-1z"/><circle cx="43" cy="37" r="2"/><path d="M46 40c0 1.7-3 6-3 6s-3-4.3-3-6s1.3-3 3-3s3 1.3 3 3z"/></g>`),
		g.Group(children),
	)
}