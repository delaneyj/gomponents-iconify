package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="m31.796 24.244l9.97 9.97l-1.415 1.414l-9.97-9.97zm-13.518.043l1.414 1.414l-9.9 9.9l-1.414-1.41z" opacity=".9"/><path fill="currentColor" d="M25 29.9c-1.5 0-3.1-.6-4.2-1.8L8.3 15.7l1.4-1.4l12.5 12.5c1.6 1.6 4.1 1.6 5.7 0l12.5-12.5l1.4 1.4l-12.6 12.5c-1.1 1.1-2.7 1.7-4.2 1.7z"/><path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v20c0 1.7-1.3 3-3 3zM11 14c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V15c0-.6-.4-1-1-1H11z"/>`),
		g.Group(children),
	)
}