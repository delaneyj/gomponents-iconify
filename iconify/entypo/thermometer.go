package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 10.123V1a1 1 0 0 0-1-1H7.799C7.247 0 7 .447 7 1v9.123A5.383 5.383 0 0 0 4.6 14.6a5.4 5.4 0 0 0 10.8 0a5.383 5.383 0 0 0-2.4-4.477zM10 17.9a3.3 3.3 0 0 1-3.3-3.3A3.29 3.29 0 0 1 9 11.471V4h2v7.471a3.29 3.29 0 0 1 2.3 3.129a3.3 3.3 0 0 1-3.3 3.3z"/>`),
		g.Group(children),
	)
}