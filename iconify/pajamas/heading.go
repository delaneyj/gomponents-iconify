package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2.75a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0v-4.5h7v4.5a.75.75 0 0 0 1.5 0V2.75a.75.75 0 0 0-1.5 0v4.5h-7v-4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}