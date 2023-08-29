package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NauseatedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#c7e755"/><g fill="#425e21"><circle cx="20.5" cy="27" r="4.5"/><circle cx="43.5" cy="27" r="4.5"/><path d="M37.4 43.2H26.6c-.7 0-.7-2.5 0-2.5h10.7c.8 0 .8 2.5.1 2.5"/><path d="M23.7 35c1 2 1.6 4.4 1.6 7c0 2.6-.6 5-1.6 7c2.1-.9 3.6-3.7 3.6-7s-1.5-6.1-3.6-7m16.6 14c-1-2-1.6-4.4-1.6-7c0-2.6.6-5 1.6-7c-2.1.9-3.6 3.7-3.6 7s1.5 6 3.6 7"/></g><path fill="#75a843" d="M25.6 15.9c-3.2 2.7-7.5 3.9-11.7 3.1c-.6-.1-1.1 2-.4 2.2c4.8.9 9.8-.5 13.5-3.6c.5-.5-1-2.1-1.4-1.7m24.5 3c-4.2.7-8.5-.4-11.7-3.1c-.4-.4-2 1.2-1.4 1.7c3.7 3.2 8.7 4.5 13.5 3.6c.7-.2.2-2.3-.4-2.2"/><g fill="#ff717f" opacity=".5"><path d="M51.7 30.3c4.9 0 8.8 4 8.8 8.8c0 4.9-4 8.8-8.8 8.8c-4.9 0-8.8-4-8.8-8.8c-.1-4.8 3.9-8.8 8.8-8.8" opacity=".8"/><circle cx="12.3" cy="39.1" r="8.8" opacity=".8"/></g>`),
		g.Group(children),
	)
}