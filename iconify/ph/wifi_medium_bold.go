package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiMediumBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 204a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm63.45-84.36a128 128 0 0 0-158.9 0a12 12 0 0 0 14.9 18.81a104 104 0 0 1 129.1 0a12 12 0 0 0 14.9-18.81Zm-32.38 35.66a80.05 80.05 0 0 0-94.14 0a12 12 0 0 0 14.14 19.4a56 56 0 0 1 65.86 0a12 12 0 1 0 14.14-19.4Z"/>`),
		g.Group(children),
	)
}