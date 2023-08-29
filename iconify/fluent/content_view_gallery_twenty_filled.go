package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 7H6v2h3V7ZM3 6a3 3 0 0 1 3-3h6v14H6a3 3 0 0 1-3-3V6Zm2 1v2a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1Zm.5 4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4ZM5 13.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0-.5.5Zm8 3.5h1a3 3 0 0 0 3-3v-1h-4v4Zm4-5V8h-4v4h4Zm0-5V6a3 3 0 0 0-3-3h-1v4h4Z"/>`),
		g.Group(children),
	)
}