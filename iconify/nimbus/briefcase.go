package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 3.75h-3.49L11.11 2a1 1 0 0 0-1-.77H6A1 1 0 0 0 5 2l-.4 1.73H1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-.98zM6.17 2.5h3.76l.29 1.25H5.88zM14.75 5v2.5H1.25V5zm-13.5 8.5V8.75H6V9a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-.25h4.75v4.75z"/>`),
		g.Group(children),
	)
}