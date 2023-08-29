package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mono(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M127 210c-44.735 0-81-36.265-81-81s36.265-81 81-81s81 36.265 81 81s-36.265 81-81 81zm1-21c34.794 0 63-27.087 63-60.5S162.794 68 128 68s-63 27.087-63 60.5S93.206 189 128 189z"/>`),
		g.Group(children),
	)
}