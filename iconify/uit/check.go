package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.353 7.574a.5.5 0 0 0-.707-.008L9.84 15.373l-3.487-3.486a.5.5 0 0 0-.707.707l3.84 3.84a.498.498 0 0 0 .707 0l8.16-8.16a.5.5 0 0 0 0-.7z"/>`),
		g.Group(children),
	)
}