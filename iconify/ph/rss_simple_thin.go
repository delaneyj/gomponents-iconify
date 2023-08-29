package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 200a4 4 0 0 1-8 0c0-81.61-66.39-148-148-148a4 4 0 0 1 0-8c86 0 156 70 156 156ZM56 116a4 4 0 0 0 0 8a76.08 76.08 0 0 1 76 76a4 4 0 0 0 8 0a84.09 84.09 0 0 0-84-84Zm4 72a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}