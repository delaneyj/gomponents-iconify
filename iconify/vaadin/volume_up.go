package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 8.5c0 2.3-.8 4.5-2 6.2l.7.8c1.5-1.9 2.4-4.4 2.4-7c0-3.1-1.2-5.9-3.2-8l-.5 1C14 3.3 15 5.8 15 8.5z"/><path fill="currentColor" d="m11.8 2.4l-.5 1C12.4 4.8 13 6.6 13 8.5c0 1.7-.5 3.2-1.3 4.6l.7.8c1.1-1.5 1.7-3.4 1.7-5.4c-.1-2.3-.9-4.4-2.3-6.1z"/><path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2zM4 5H0v6h4l5 4V1z"/>`),
		g.Group(children),
	)
}