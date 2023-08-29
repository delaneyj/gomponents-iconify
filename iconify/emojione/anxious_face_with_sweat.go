package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnxiousFaceWithSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#664e27" d="M40.3 48.1c0 4.6-3.7 8.3-8.3 8.3c-4.6 0-8.3-3.7-8.3-8.3c0-4.6 3.7-8.3 8.3-8.3c4.6 0 8.3 3.7 8.3 8.3"/><path fill="#fff" d="M26.2 44.8c1.2-2 3.3-3.4 5.8-3.4s4.7 1.3 5.8 3.4H26.2"/><g fill="#664e27"><circle cx="43.5" cy="33" r="4.5"/><circle cx="20.5" cy="33" r="4.5"/></g><path fill="#917524" d="M25.6 17.9c-3.2 2.7-7.5 3.9-11.7 3.1c-.6-.1-1.1 2-.4 2.2c4.8.9 9.8-.5 13.5-3.6c.5-.5-1-2.1-1.4-1.7m24.5 3c-4.2.7-8.5-.4-11.7-3.1c-.4-.4-2 1.2-1.4 1.7c3.7 3.2 8.7 4.5 13.5 3.6c.7-.2.2-2.3-.4-2.2"/><path fill="#65b1ef" d="M62 18.5c0 9.4-12.7 9.4-12.7 0c0-6.9 6.4-13.5 6.4-13.5S62 11.7 62 18.5"/>`),
		g.Group(children),
	)
}