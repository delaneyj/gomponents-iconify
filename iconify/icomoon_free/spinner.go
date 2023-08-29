package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 2zm4.243 1.757a2 2 0 1 1 3.999-.001a2 2 0 0 1-3.999.001zM13 8a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm-1.757 4.243a1 1 0 1 1 2 0a1 1 0 0 1-2 0zM7 14a1 1 0 0 1 2 0a1 1 0 0 1-2 0zm-4.243-1.757a1 1 0 0 1 2 0a1 1 0 0 1-2 0zm-.5-8.486a1.5 1.5 0 0 1 3 0a1.5 1.5 0 0 1-3 0zM.875 8a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0z"/>`),
		g.Group(children),
	)
}