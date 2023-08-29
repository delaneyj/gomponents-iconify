package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaLibraryFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M13 15v8l7-4l-7-4z"/><path fill="currentColor" d="M26 10H6a2.002 2.002 0 0 0-2 2v14a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V12a2.002 2.002 0 0 0-2-2zM13 23v-8l7 4zM6 6h20v2H6zm2-4h16v2H8z"/>`),
		g.Group(children),
	)
}