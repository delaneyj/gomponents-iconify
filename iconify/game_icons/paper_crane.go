package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperCrane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m174.2 108.9l24.3 157.3l31.8-15.6zM383 142l28.8 40.6l-38.3-30.4l-2.8 15.3l60.5 33.2l-35.8-61.3zm-18.1 5.5l-13.8 7.2l-41.2 143.4l-3.2 105l11.1-16.6zm9.8 38.3l-37.1 188.3l12.8 7.9zm115.8 13.5l-107.4 21.3l-17.2 125l24.4-73l33.3-6.4zM315 234.7l-21.7 4.8l-3.7 9.4l20.1 1.7zm-38.4 23.4c-37.4-.3-69.1 25.4-69.1 25.4l69.8 14.4l18.2 67.7l-1.2-68.4l11.6-32.9c-9.9-4.3-19.8-6.1-29.3-6.2zM170.2 290l-38.4 8.8l-106.33 68.6L218.5 340l31-16.9l35.3 63.4l-20.2-78.1zm78.3 58l-41.1 4.5l4 20.6h35.9l33.4 27.1z"/>`),
		g.Group(children),
	)
}