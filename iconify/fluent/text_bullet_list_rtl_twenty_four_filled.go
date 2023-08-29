package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRtlTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.504 16.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-4 .5H3a1 1 0 0 0-.117 1.993L3 19h13.503a1 1 0 0 0 .117-1.993L16.503 17Zm4-6.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-4 .5H3a1 1 0 0 0-.117 1.993L3 13h13.503a1 1 0 0 0 .117-1.993L16.503 11Zm4-6.492a1.5 1.5 0 1 0 0 2.999a1.5 1.5 0 0 0 0-3Zm-4 .493H3a1 1 0 0 0-.117 1.993L3 7.001h13.503a1 1 0 0 0 .117-1.993L16.503 5Z"/>`),
		g.Group(children),
	)
}