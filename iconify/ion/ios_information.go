package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosInformation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M235.4 176c0-11.4 9.3-19.9 20.5-19.9 11.4 0 20.7 8.5 20.7 19.9s-9.3 20-20.7 20c-11.2.1-20.5-8.6-20.5-20zm1.4 35.8H275v144.1h-38.2V211.8z" fill="currentColor"/>`),
		g.Group(children),
	)
}