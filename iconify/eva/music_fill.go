package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMusicFill0"><g id="evaMusicFill1"><path id="evaMusicFill2" fill="currentColor" d="M19 15V4a1 1 0 0 0-.38-.78a1 1 0 0 0-.84-.2l-9 2A1 1 0 0 0 8 6v8.34a3.49 3.49 0 1 0 2 3.18a4.36 4.36 0 0 0 0-.52V6.8l7-1.55v7.09a3.49 3.49 0 1 0 2 3.17a4.57 4.57 0 0 0 0-.51Z"/></g></g>`),
		g.Group(children),
	)
}