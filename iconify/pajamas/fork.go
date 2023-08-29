package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.25 2.386a2.501 2.501 0 1 0-1.5 0v.364a2.5 2.5 0 0 0 2.5 2.5a1 1 0 0 1 1 1v.364a2.501 2.501 0 1 0 1.5 0V9.75a1 1 0 0 1 1-1a2.5 2.5 0 0 0 2.5-2.5v-.364a2.501 2.501 0 1 0-1.5 0v.364a1 1 0 0 1-1 1c-.681 0-1.3.273-1.75.715a2.492 2.492 0 0 0-1.75-.715a1 1 0 0 1-1-1v-.364ZM11.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-3.5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}