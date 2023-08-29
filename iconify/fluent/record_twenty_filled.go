package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm8 4.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/>`),
		g.Group(children),
	)
}