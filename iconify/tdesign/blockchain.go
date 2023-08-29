package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockchain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h7v2.5h6V2h7v7h-2.5v6H22v7h-7v-2.5H9V22H2v-7h2.5V9H2V2Zm5 5V4H4v3h3Zm-.5 2v6H9v2.5h6V15h2.5V9H15V6.5H9V9H6.5ZM17 17v3h3v-3h-3ZM7 17H4v3h3v-3ZM17 4v3h3V4h-3Z"/>`),
		g.Group(children),
	)
}