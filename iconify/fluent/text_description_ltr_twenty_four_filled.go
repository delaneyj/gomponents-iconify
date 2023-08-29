package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDescriptionLtrTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 17h12a1 1 0 0 1 .117 1.993L15 19H3a1 1 0 0 1-.117-1.993L3 17h12H3Zm0-4h18a1 1 0 0 1 .117 1.993L21 15H3a1 1 0 0 1-.117-1.993L3 13h18H3Zm0-4h18a1 1 0 0 1 .117 1.993L21 11H3a1 1 0 0 1-.117-1.993L3 9h18H3Zm0-4h18a1 1 0 0 1 .117 1.993L21 7H3a1 1 0 0 1-.117-1.993L3 5h18H3Z"/>`),
		g.Group(children),
	)
}