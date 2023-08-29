package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopBottomThickTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4.502a1.5 1.5 0 0 0-1.396.949a.75.75 0 1 1-1.413-.492h-.005A3.001 3.001 0 0 1 6 3.002h8c1.29 0 2.39.814 2.814 1.957h-.003a.75.75 0 0 1-1.381.586l-.004-.01A1.5 1.5 0 0 0 14 4.501H6ZM3.75 12a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v2.5a.75.75 0 0 1-.75.75ZM14 15.5h-2.711a.825.825 0 0 0-.039 0H6a1.5 1.5 0 0 1-1.396-.948a.75.75 0 1 0-1.413.492h-.005A3.001 3.001 0 0 0 6 17.002h8c1.29 0 2.39-.814 2.814-1.957h-.003a.75.75 0 0 0-1.381-.586l-.004.01A1.5 1.5 0 0 1 14 15.501Zm1.5-4.25a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5Z"/>`),
		g.Group(children),
	)
}