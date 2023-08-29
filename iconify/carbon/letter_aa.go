package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterAa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 13h-5v2h5v2h-4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6v-8a2 2 0 0 0-2-2zm0 8h-4v-2h4zM13 9H9a2 2 0 0 0-2 2v12h2v-5h4v5h2V11a2 2 0 0 0-2-2zm-4 7v-5h4v5z"/>`),
		g.Group(children),
	)
}