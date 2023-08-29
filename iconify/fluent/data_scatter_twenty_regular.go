package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScatterTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.5a.5.5 0 0 0-1 0v15a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 0-1H3V2.5Zm3 5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0ZM7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Zm7-1a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM12 5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Zm-1 7a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}