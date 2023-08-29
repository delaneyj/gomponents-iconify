package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m259.4 40.33l-93 60.37L64.03 58.34L92.67 165.4l-71.89 84.3l110.62 5.8l58.1 94.5l32.4-84.6l208.7 206.8l25.4-25.4l-.1-.1l32.6-32.5l-25.4-25.4l-32.6 32.5l-18.4-18.4l21.2-21.2l-25.4-25.4l-21.2 21.2l-17-17l55.1-55.2l-25.4-25.4l-55.1 55.2l-95.2-93.4l87.8-21.1l-86.1-69.8l8.6-110.47z"/>`),
		g.Group(children),
	)
}