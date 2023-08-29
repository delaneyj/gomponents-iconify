package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStyleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Zm7.5 0a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Zm7.5 0a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Zm3.502 6.503h-18.5l-.101.007a.75.75 0 0 0 .101 1.493h18.5l.102-.007a.75.75 0 0 0-.102-1.493ZM3.25 17a1.25 1.25 0 0 0 0 2.5h17.5a1.25 1.25 0 1 0 0-2.5H3.25Z"/>`),
		g.Group(children),
	)
}