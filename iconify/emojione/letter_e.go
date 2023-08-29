package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M42.3 22.6H26.9v6.2H41v5H26.9v7.5H43v5.2H21v-29h21.3v5.1z"/>`),
		g.Group(children),
	)
}