package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drumpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32 64.001C32 46.328 46.333 32 64.001 32H192c17.672 0 32 14.333 32 32.001V192c0 17.672-14.333 32-32.001 32H64c-17.672 0-32-14.333-32-32.001V64z"/>`),
		g.Group(children),
	)
}