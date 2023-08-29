package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M20.3 46.5v-29h6v11.1h11.3V17.5h6v29h-6v-13H26.4v13h-6.1z"/>`),
		g.Group(children),
	)
}