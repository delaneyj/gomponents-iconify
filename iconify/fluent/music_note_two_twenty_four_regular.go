package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.698 2.148A.75.75 0 0 1 20 2.75v13.5a.764.764 0 0 1-.004.079a3.5 3.5 0 1 1-1.496-2.702V7.758l-8.5 2.55v7.942a.756.756 0 0 1-.004.079A3.5 3.5 0 1 1 8.5 15.627V5.75a.75.75 0 0 1 .534-.718l10-3a.75.75 0 0 1 .664.116ZM10 8.742l8.5-2.55V3.758L10 6.308v2.434ZM6.5 16.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm8 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		g.Group(children),
	)
}