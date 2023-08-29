package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M19.5 42.1L20 20l22.1-.5l-7.5 7.5l13.3 13.2l-7.3 7.3l-13.2-13.3z"/>`),
		g.Group(children),
	)
}