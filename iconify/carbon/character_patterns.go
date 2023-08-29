package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterPatterns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 26v-8.172l-3.586 3.586L1 20l6-6l6 6l-1.414 1.414L8 17.828V26h10v2H8a2.002 2.002 0 0 1-2-2zm24-4v-2h-8v2h5.5L22 28v2h8v-2h-5.493L30 22zM26 6v8.172l3.586-3.586L31 12l-6 6l-6-6l1.414-1.414L24 14.172V6H14V4h10a2.002 2.002 0 0 1 2 2zM8 2H3v2h5v2H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6V4a2.002 2.002 0 0 0-2-2zm0 8H4V8h4z"/>`),
		g.Group(children),
	)
}