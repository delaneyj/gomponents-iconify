package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineCharacterUTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 3.5a.5.5 0 0 0-1 0V10a4 4 0 0 0 8 0V3.5a.5.5 0 0 0-1 0V10a3 3 0 1 1-6 0V3.5ZM5.5 16a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1h-9Z"/>`),
		g.Group(children),
	)
}