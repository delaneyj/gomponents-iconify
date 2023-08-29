package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymDifference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M65 10.875c-8.937 0-17.097 3.371-23.29 8.902C57.81 22.917 70 37.12 70 54.125c0 10.364-4.53 19.684-11.71 26.098c2.172.423 4.415.652 6.71.652c19.3 0 35-15.7 35-35s-15.7-35-35-35zm-6.71 69.348C42.19 77.083 30 62.88 30 45.875c0-10.364 4.53-19.684 11.71-26.098a35.026 35.026 0 0 0-6.71-.652c-19.3 0-35 15.7-35 35s15.7 35 35 35c8.937 0 17.097-3.371 23.29-8.902z" color="currentColor"/>`),
		g.Group(children),
	)
}