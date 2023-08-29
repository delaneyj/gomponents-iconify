package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm0 4.75A1.25 1.25 0 1 1 9 8a1.25 1.25 0 0 1 2.5 0Zm-5.75 6a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm4.5 0a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm-4.5-4.75a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm0-4.75a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}