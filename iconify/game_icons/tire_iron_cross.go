package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TireIronCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M52.47 16.94L16.95 52.45L71.3 106.6h22.59L215 227.7v56.5L93.84 405.4H71.28l-54.26 54.2l35.34 35.5l54.24-54.4v-22.5L227.8 297h56.4l121.2 121.2v22.5l54.2 54.3l35.4-35.4l-54.3-54.2h-22.5L297 284.2v-56.5l121.1-121.1h22.6L495 52.36L459.7 17l-54.3 54.25v22.57L284.2 215h-56.4L106.6 93.86V71.28L52.47 16.94z"/>`),
		g.Group(children),
	)
}