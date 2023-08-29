package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squaretwitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-83-768q-6 3-18 11.5t-19.5 13t-18.5 9.5t-21 7q-37-41-91-41q-117 0-117 98v60q-161-9-247-119q-25 26-25 57q0 66 49 100q-7 0-17.5 1t-17 0t-14.5-5q0 46 24.5 76.5t67.5 39.5q-10 12-28 12q-16 0-28-9q0 39 37.5 60.5t84.5 22.5q-18 27-52.5 40.5t-73.5 13.5q-14 0-38.5-7t-25.5-7q16 32 65.5 55t125.5 23q67 0 125-23.5t99-62.5t70.5-89t44-103.5t14.5-105.5q0-2 12-8.5t28-17t24-23.5q-54 0-72 2q35-21 53-81z"/>`),
		g.Group(children),
	)
}