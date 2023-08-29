package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MantelpieceClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#594640" d="M60 40c-5 0-9-8-9-13c0-10-9-18-19-18s-19 8-19 18c0 5-4 13-9 13H0v13h64V40h-4z"/><circle cx="32" cy="28" r="15" fill="#fed0ac"/><g fill="#333"><circle cx="32" cy="28" r="2"/><path d="M40.1 30c.7 0 2.9-2 2.9-2s-2.2-2-2.9-2c-.6 0-1.1.6-1.3 1.4H31v1.2h7.8c.2.8.7 1.4 1.3 1.4"/><path d="m32 13.7l-1 14h2z"/></g><path fill="#d3976e" d="M3 53h8v2H3zm50 0h8v2h-8z"/>`),
		g.Group(children),
	)
}