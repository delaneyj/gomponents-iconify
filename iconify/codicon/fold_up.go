package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1 7.4l.7.7l6-6l6 6l.7-.7L8.1 1h-.7L1 7.4zm0 6l.7.7l6-6l6 6l.7-.7L8.1 7h-.7L1 13.4z"/>`),
		g.Group(children),
	)
}