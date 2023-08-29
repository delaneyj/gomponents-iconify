package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#212121" d="M212 32.2c-17.56 0-31.79 14.23-31.79 31.79S194.44 95.8 212 95.8s31.79-14.25 31.79-31.82S229.56 32.2 212 32.2zM64 24.3c-21.92 0-39.69 17.78-39.69 39.71S42.08 103.7 64 103.7s39.69-17.77 39.69-39.69S85.92 24.3 64 24.3z"/>`),
		g.Group(children),
	)
}