package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 2H10a2 2 0 1 0-4 0H1.5a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-13a.5.5 0 0 0-.5-.5zM8 1a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm6 14H2V3h2v1.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V3h2v12z"/><path fill="currentColor" d="M7 13.414L3.793 9.707l.914-.914L7 10.586l4.293-3.793l.914.914z"/>`),
		g.Group(children),
	)
}