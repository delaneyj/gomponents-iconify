package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.78 2.044a.73.73 0 0 1 .72.741v14.473a.73.73 0 0 1-.72.742a.731.731 0 0 1-.72-.742V2.785a.73.73 0 0 1 .72-.741ZM5.22 2a.73.73 0 0 1 .72.742v14.473a.73.73 0 0 1-.72.741a.731.731 0 0 1-.72-.741V2.742A.73.73 0 0 1 5.22 2Z"/>`),
		g.Group(children),
	)
}