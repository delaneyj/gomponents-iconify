package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pipeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.272 3.864a4.5 4.5 0 1 0-.1 8.314a.75.75 0 1 0-.557-1.393a3 3 0 1 1 .066-5.543a.75.75 0 1 0 .59-1.378ZM11.5 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 1.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM6 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}