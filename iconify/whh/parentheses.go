package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parentheses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 1024h-64q128-160 128-512T768 0h64q91 85 141.5 216.5T1024 512t-50.5 295.5T832 1024zM192 0h64Q128 160 128 512t128 512h-64q-91-85-141.5-216.5T0 512t50.5-295.5T192 0z"/>`),
		g.Group(children),
	)
}