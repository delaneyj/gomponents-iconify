package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.43 7.686L12 18.4l8.57-10.713a15.015 15.015 0 0 0-17.14 0Zm-2.051-.961c6.192-4.967 15.05-4.967 21.243 0l.779.625l-11.4 14.25L.6 7.35l.779-.625Z"/>`),
		g.Group(children),
	)
}