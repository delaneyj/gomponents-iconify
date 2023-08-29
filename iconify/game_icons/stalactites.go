package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stalactites(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M18 18v109.8l41.97 274L109.4 79.5l31.5 161.9l19.5-78l27.8 161.2l55.4-272.9l54 172.4l14.2-35.2l60.8 289.4l59.2-320.1l24.9 57.3c-2.9 6-7.1 14-5.4 21.1c1.1 4.8 4.7 11.2 9.6 11.1c4.7-.1 7.7-6.5 8.5-11.1c1.4-7.6-3.6-16.2-6.8-22.1l31.4-74.8V18zm170.1 329.2s-10.4 17.8-8.5 27c1 4.8 4.7 11.2 9.6 11.1c4.7-.1 7.8-6.5 8.5-11.1c1.5-9.4-9.6-27-9.6-27z"/>`),
		g.Group(children),
	)
}