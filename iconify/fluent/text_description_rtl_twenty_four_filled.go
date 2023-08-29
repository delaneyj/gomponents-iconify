package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDescriptionRtlTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 17H9a1 1 0 0 0-.117 1.993L9 19h12a1 1 0 0 0 .117-1.993L21 17H9h12Zm0-4H3a1 1 0 0 0-.117 1.993L3 15h18a1 1 0 0 0 .117-1.993L21 13H3h18Zm0-4H3a1 1 0 0 0-.117 1.993L3 11h18a1 1 0 0 0 .117-1.993L21 9H3h18Zm0-4H3a1 1 0 0 0-.117 1.993L3 7h18a1 1 0 0 0 .117-1.993L21 5H3h18Z"/>`),
		g.Group(children),
	)
}