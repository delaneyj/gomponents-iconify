package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Church(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.447 14.105L18 12.382V12a1 1 0 0 0-.485-.857L13 8.434V6h2V4h-2V2h-2v2H9v2h2v2.434l-4.515 2.709A1 1 0 0 0 6 12v.382l-3.447 1.724A.998.998 0 0 0 2 15v6a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-6c0-.379-.214-.725-.553-.895zM4 15.618l2-1V20H4v-4.382zM12 15a2 2 0 0 0-2 2v3H8v-7.434l4-2.4l4 2.4V20h-2v-3a2 2 0 0 0-2-2zm8 5h-2v-5.382l2 1V20z"/>`),
		g.Group(children),
	)
}