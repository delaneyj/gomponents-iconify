package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="255.75" cy="56" r="56" fill="currentColor"/><path fill="currentColor" d="M310.28 191.4h.05l7.66-2.3l36.79 122.6l46-13.8l-16.21-54.16c0-.12 0-.24-.07-.36l-16.84-56.12l-4.71-15.74l-.9-3H362l-2.51-8.45a44.84 44.84 0 0 0-43-32.08H195.24a44.84 44.84 0 0 0-43 32.08l-2.51 8.45h-.06l-.9 3l-4.71 15.74l-16.84 56.12c0 .12 0 .24-.07.36l-16.21 54.16l46 13.8l36.76-122.6l7.54 2.26L148.25 368h51.5v144h52V368h8v144h52V368h51.51Z"/>`),
		g.Group(children),
	)
}