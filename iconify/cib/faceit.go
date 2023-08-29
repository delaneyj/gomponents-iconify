package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faceit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 3.599c0-.13-.135-.266-.135-.266c-.13 0-.13 0-.266.135c-2.667 4.13-5.464 8.266-8.13 12.531H.266c-.266 0-.401.401-.13.531c9.599 3.604 23.599 9.068 31.333 12.135c.266.135.531-.135.531-.266z"/>`),
		g.Group(children),
	)
}