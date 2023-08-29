package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopBottomTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4.502a1.5 1.5 0 0 0-1.396.949a.75.75 0 1 1-1.413-.492h-.005A3.001 3.001 0 0 1 6 3.002h8c1.29 0 2.39.814 2.814 1.957h-.003a.75.75 0 0 1-1.381.586l-.004-.01A1.5 1.5 0 0 0 14 4.501H6ZM3 11.25a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5Zm12.5 0a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5ZM4.604 14.553A1.5 1.5 0 0 0 6 15.502h8a1.5 1.5 0 0 0 1.426-1.033l.003-.01a.75.75 0 0 1 1.382.586h.003A3.001 3.001 0 0 1 14 17.002H6c-1.29 0-2.39-.814-2.814-1.957h.005a.75.75 0 1 1 1.413-.492Z"/>`),
		g.Group(children),
	)
}