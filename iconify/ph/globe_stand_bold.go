package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeStandBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 172a76 76 0 1 0-76-76a76.08 76.08 0 0 0 76 76Zm0-128a52 52 0 1 1-52 52a52.06 52.06 0 0 1 52-52Zm12 167.38V220h20a12 12 0 0 1 0 24H96a12 12 0 0 1 0-24h20v-8.62A116 116 0 0 1 12 97.12a115.3 115.3 0 0 1 32.29-81.43A12 12 0 1 1 61.6 32.31a92 92 0 0 0 130.09 130.08a12 12 0 1 1 16.62 17.31A115.12 115.12 0 0 1 140 211.38Z"/>`),
		g.Group(children),
	)
}