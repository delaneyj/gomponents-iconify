package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YCombinator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m809 876l266-499H963L806 689q-24 48-44 92l-42-92l-155-312H445l263 493v324h101V876zM1536 0v1536H0V0h1536z"/>`),
		g.Group(children),
	)
}