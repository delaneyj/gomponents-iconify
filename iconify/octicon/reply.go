package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M6 3.5c3.92.44 8 3.125 8 10c-2.312-5.062-4.75-6-8-6V11L.5 5.5L6 0v3.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}