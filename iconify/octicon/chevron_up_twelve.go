package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 4c-.2 0-.4.1-.5.2L2.2 7.5c-.3.3-.3.8 0 1.1c.3.3.8.3 1.1 0L6 5.9l2.7 2.7c.3.3.8.3 1.1 0c.3-.3.3-.8 0-1.1L6.6 4.3C6.4 4.1 6.2 4 6 4Z"/>`),
		g.Group(children),
	)
}