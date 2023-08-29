package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterFraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30h-8v-6a2.002 2.002 0 0 1 2-2h4v-4h-6v-2h6a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2h-4v4h6ZM7 23.586L22.586 8L24 9.414L8.414 25zM4.5 15.5v-1h3v-11h-3v-1h4v12h3v1h-7z"/><path fill="currentColor" d="M8 3v12V3m1-1H4v2h3v10H4v2h8v-2H9V2Z"/>`),
		g.Group(children),
	)
}