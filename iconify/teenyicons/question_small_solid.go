package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 6a2 2 0 0 1 2-2h.6c1.05 0 1.9.85 1.9 1.9V6a2 2 0 0 1-2 2v1H7V7h1a1 1 0 0 0 1-1v-.1a.9.9 0 0 0-.9-.9h-.6a1 1 0 0 0-1 1h-1ZM8 10v1H7v-1h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}