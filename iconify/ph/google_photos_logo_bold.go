package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotosLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 116h-33.36A76 76 0 0 0 128 12a12 12 0 0 0-12 12v33.36A76 76 0 0 0 12 128a12 12 0 0 0 12 12h33.36A76 76 0 0 0 128 244a12 12 0 0 0 12-12v-33.36A76 76 0 0 0 244 128a12 12 0 0 0-12-12Zm-52-28a51.38 51.38 0 0 1-8.18 28H140V37.4A52.09 52.09 0 0 1 180 88ZM88 76a51.38 51.38 0 0 1 28 8.18V116H37.4A52.09 52.09 0 0 1 88 76Zm-12 92a51.38 51.38 0 0 1 8.18-28H116v78.6A52.09 52.09 0 0 1 76 168Zm92 12a51.38 51.38 0 0 1-28-8.18V140h78.6a52.09 52.09 0 0 1-50.6 40Z"/>`),
		g.Group(children),
	)
}