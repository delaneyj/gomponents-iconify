package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Out(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.5 8c.3 0 .5.2.5.5v2c0 .3-.2.5-.5.5s-.5-.2-.5-.5v-2c0-.3.2-.5.5-.5zm0-1C2.7 7 2 7.7 2 8.5v2c0 .8.7 1.5 1.5 1.5S5 11.3 5 10.5v-2C5 7.7 4.3 7 3.5 7zM8 7v3.5c0 .3-.2.5-.5.5s-.5-.2-.5-.5V7H6v3.5c0 .8.7 1.5 1.5 1.5S9 11.3 9 10.5V7H8zm5 0h-3v1h1v4h1V8h1z"/><path fill="currentColor" d="M15 6V5h-2.4L8.9 2c.1-.2.1-.3.1-.5C9 .7 8.3 0 7.5 0S6 .7 6 1.5c0 .2 0 .3.1.5L2.4 5H0v9h1v1h15V6h-1zM6.7 2.8c.3.1.5.2.8.2s.5-.1.8-.2L11 5H4l2.7-2.2zM14 13H1V6h13v7z"/>`),
		g.Group(children),
	)
}