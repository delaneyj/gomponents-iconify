package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.31 66.46l-96-40a6.06 6.06 0 0 0-4.62 0l-96 40a6 6 0 0 0-3.64 6.33l16 120a6 6 0 0 0 3.27 4.58l80 40a6 6 0 0 0 5.36 0l80-40a6 6 0 0 0 3.27-4.58l16-120a6 6 0 0 0-3.64-6.33Zm-23.84 121.6L128 225.29l-74.47-37.23l-15-112.29L128 38.5l89.44 37.27Zm-79.72-103l-40 72a6 6 0 0 0 10.5 5.82L104.86 142h46.28l11.61 20.91A6 6 0 0 0 168 166a5.88 5.88 0 0 0 2.9-.76a6 6 0 0 0 2.34-8.15l-40-72a6 6 0 0 0-10.5 0ZM144.47 130h-32.94L128 100.35Z"/>`),
		g.Group(children),
	)
}