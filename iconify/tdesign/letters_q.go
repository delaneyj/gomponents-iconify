package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersQ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6v12h6V6H9Zm8 12h1.5v2h-10A1.5 1.5 0 0 1 7 18.5v-13A1.5 1.5 0 0 1 8.5 4h7A1.5 1.5 0 0 1 17 5.5V18Z"/>`),
		g.Group(children),
	)
}