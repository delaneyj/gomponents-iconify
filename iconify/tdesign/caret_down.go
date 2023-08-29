package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.914L21.414 8.5H2.586L12 17.914ZM7.414 10.5h9.172L12 15.086L7.414 10.5Z"/>`),
		g.Group(children),
	)
}