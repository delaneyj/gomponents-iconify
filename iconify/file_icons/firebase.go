package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firebase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M323.669 114.555c-1.516-9.335-11.112-13.198-17.794-6.505l-62.225 62.165l-47.868-91.038c-4.473-8.515-15.34-9.407-19.765-.868l-27.129 51.66L83.521 6.332C78.127-3.8 64.427-1.217 62.657 10.125L0 410.825l143.038-269.622l46.368 85.666L.054 413.574l168.103 93.96a35.202 35.202 0 0 0 34.306.008l169.653-94.617l-48.447-298.37z"/>`),
		g.Group(children),
	)
}