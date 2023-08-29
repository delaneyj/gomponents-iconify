package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6 15.321l9.014-4.883L6 4.804v10.517Zm9.49-4.003a1 1 0 0 0 .054-1.728L6.53 3.956A1 1 0 0 0 5 4.804v10.517a1 1 0 0 0 1.476.88l9.015-4.883Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}