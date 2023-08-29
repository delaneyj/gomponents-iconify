package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceUnhappy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13ZM8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16Zm0-7a3 3 0 0 0-3 3h6a3 3 0 0 0-3-3ZM6.002 5a1.5 1.5 0 0 0-1.3.747a.5.5 0 1 0 .866.502a.5.5 0 0 1 .865 0a.5.5 0 0 0 .866-.5A1.5 1.5 0 0 0 6.002 5Zm3.25.2a1.5 1.5 0 0 1 2.047.55a.5.5 0 1 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}