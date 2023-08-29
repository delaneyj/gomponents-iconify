package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFourThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 176a4 4 0 0 1-4 4h-12v28a4 4 0 0 1-8 0v-28h-44a4 4 0 0 1-3.79-5.26l24-72a4 4 0 1 1 7.58 2.52L189.55 172H228v-28a4 4 0 0 1 8 0v28h12a4 4 0 0 1 4 4ZM144 52a4 4 0 0 0-4 4v56H44V56a4 4 0 0 0-8 0v120a4 4 0 0 0 8 0v-56h96v56a4 4 0 0 0 8 0V56a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}