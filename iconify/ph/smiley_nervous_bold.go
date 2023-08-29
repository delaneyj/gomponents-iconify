package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyNervousBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184.49 176.49a12 12 0 0 1-17 0L160 169l-7.51 7.52a12 12 0 0 1-17 0L128 169l-7.51 7.52a12 12 0 0 1-17 0L96 169l-7.51 7.52a12 12 0 0 1-17-17l16-16a12 12 0 0 1 17 0L112 151l7.51-7.52a12 12 0 0 1 17 0L144 151l7.51-7.52a12 12 0 0 1 17 0l16 16a12 12 0 0 1-.02 17.01ZM236 128A108 108 0 1 1 128 20a108.12 108.12 0 0 1 108 108Zm-24 0a84 84 0 1 0-84 84a84.09 84.09 0 0 0 84-84Zm-120-4a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm72 0a16 16 0 1 0-16-16a16 16 0 0 0 16 16Z"/>`),
		g.Group(children),
	)
}