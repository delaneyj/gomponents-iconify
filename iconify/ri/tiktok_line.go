package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TiktokLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2v6.414a6.85 6.85 0 1 0 5.6 6.736v-3.736a8.603 8.603 0 0 0 3.4.686h1V6.5h-1c-1.903 0-3.4-1.537-3.4-3.5V2H11Zm2 2h1.688c.394 2.22 2.08 3.996 4.312 4.41v1.618c-1.038-.152-1.975-.542-2.843-1.123L14.6 7.863v7.287a4.85 4.85 0 1 1-4.6-4.844v1.604a3.25 3.25 0 1 0 3 3.24V4ZM8.5 15.15a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z"/>`),
		g.Group(children),
	)
}