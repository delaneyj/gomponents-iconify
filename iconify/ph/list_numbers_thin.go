package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbersThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4H104a4 4 0 0 1 0-8h112a4 4 0 0 1 4 4ZM104 68h112a4 4 0 0 0 0-8H104a4 4 0 0 0 0 8Zm112 120H104a4 4 0 0 0 0 8h112a4 4 0 0 0 0-8ZM41.79 51.58L52 46.47V104a4 4 0 0 0 8 0V40a4 4 0 0 0-5.79-3.58l-16 8a4 4 0 1 0 3.58 7.16ZM72 204H48l23.85-31.92a19.54 19.54 0 0 0 4-14.8a19.76 19.76 0 0 0-8-13.28a20.84 20.84 0 0 0-28.59 3.92a19.85 19.85 0 0 0-3 5.38a4 4 0 0 0 7.5 2.7a12.1 12.1 0 0 1 1.78-3.22a12.78 12.78 0 0 1 17.54-2.37a11.85 11.85 0 0 1 4.81 7.94a11.65 11.65 0 0 1-2.41 8.85L36.8 205.61A4 4 0 0 0 40 212h32a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}