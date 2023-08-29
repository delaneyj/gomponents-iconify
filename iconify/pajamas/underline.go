package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 2.75a.75.75 0 0 0-1.5 0V7a4 4 0 1 0 8 0V2.75a.75.75 0 0 0-1.5 0V7a2.5 2.5 0 0 1-5 0V2.75ZM3.75 12.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}