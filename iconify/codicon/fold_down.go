package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.207 1.707L13.5 1l-6 6l-6-6l-.707.707l6.353 6.354h.708l6.353-6.354zm0 6L13.5 7l-6 6l-6-6l-.707.707l6.353 6.354h.708l6.353-6.354z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}