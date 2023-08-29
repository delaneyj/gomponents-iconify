package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v3.5a.5.5 0 0 0 .5.5H7M3 11V7.5a.5.5 0 0 1 .5-.5H7m18 0h3.5a.5.5 0 0 1 .5.5V11m0 10v3.5a.5.5 0 0 1-.5.5H25"/>`),
		g.Group(children),
	)
}