package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BangingGavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M155 18L80.47 38.98l15.9 55.79L283.5 41.95L276.8 18zm105.5 48l-131 37l61.1 216.4l130.9-37zm53.3 52.9l-19.4 5.5l24 85l19.5-5.5zm-182.7 51.6l-19.5 5.5l24 85l19.5-5.4zm-30.9 27.6L18 221.3v54.3l96.5-27.4zm287.4 19.7l-55.7 34.7l6.5 24.6l28.4-8l24.4 89.6L171 421.3c-8.4-30-16.9-60-25.3-90l27.3-7.7l-6.3-22.7l-70.53-3.8L137 336.7L28.26 385.6s117.34 4.1 114.34 4.6c-3.1.5-31.3 84.4-31.3 84.4l88-45.2l22.9 64.3l70.6-76.4l94.4 49.7l-24.7-70.9l113.5-5.6l-77.7-53.7l94.6-66.3l-113.4 3.5zM354 290.7l-187.3 52.9l15.7 55.9l187.2-52.9z"/>`),
		g.Group(children),
	)
}