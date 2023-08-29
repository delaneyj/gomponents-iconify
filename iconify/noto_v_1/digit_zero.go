package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fff" d="M62.79 25C42.87 25 30 40.47 30 64.42c0 23.95 12.87 39.43 32.79 39.43s32.79-15.47 32.79-39.43C95.59 40.47 82.71 25 62.79 25zm0 60.21c-9.53 0-12.91-11.2-12.91-20.79c0-9.58 3.38-20.78 12.91-20.78s12.92 11.2 12.92 20.78c0 9.59-3.39 20.79-12.92 20.79z"/>`),
		g.Group(children),
	)
}