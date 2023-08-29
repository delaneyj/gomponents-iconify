package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.35 12.79a.72.72 0 0 1-.49-.19a7.24 7.24 0 0 0-9.66 0a.75.75 0 0 1-1.06-.06a.76.76 0 0 1 .07-1.06a8.7 8.7 0 0 1 11.64 0a.76.76 0 0 1 .07 1.06a.79.79 0 0 1-.57.25Z"/><path fill="currentColor" d="M20 10a.76.76 0 0 1-.51-.2a10.87 10.87 0 0 0-15 0a.75.75 0 1 1-1-1.1a12.36 12.36 0 0 1 17 0A.75.75 0 0 1 20 10ZM9.38 15.64a.75.75 0 0 1-.6-.3a.74.74 0 0 1 .15-1.05a5.06 5.06 0 0 1 6.15 0a.75.75 0 0 1 .15 1.05a.76.76 0 0 1-1.05.15a3.59 3.59 0 0 0-4.35 0a.78.78 0 0 1-.45.15ZM12 18.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}