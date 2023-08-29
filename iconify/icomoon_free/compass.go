package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.5 16a.5.5 0 0 1-.5-.5V8H.5a.5.5 0 0 1-.211-.953l15-7a.5.5 0 0 1 .665.665l-7 15a.5.5 0 0 1-.453.289zM2.754 7H8.5a.5.5 0 0 1 .5.5v5.746l5.465-11.712L2.753 6.999z"/>`),
		g.Group(children),
	)
}