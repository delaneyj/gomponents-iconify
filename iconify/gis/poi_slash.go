package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoiSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m68.913 48.908l-.048.126c.015-.038.027-.077.042-.115l.006-.011z" color="currentColor"/><path fill="currentColor" d="M50.002 0C30.764 0 15 15.719 15 34.902c0 7.433 2.374 14.339 6.393 20.018l2.72 4.703L36.28 47.457a18.442 18.442 0 0 1-4.88-12.555c0-10.33 8.243-18.548 18.602-18.548c4.87 0 9.256 1.83 12.553 4.828L74.082 9.654C67.794 3.685 59.308 0 50.002 0zm34.824 31.438L36.033 80.23l9.697 16.764c3.409 4.453 5.674 3.608 8.508-.234l26.846-45.684c.542-.98.967-2.025 1.338-3.092a34.446 34.446 0 0 0 2.404-16.547z"/><path fill="currentColor" d="M87.123 2.27L97.73 12.878L12.877 97.73L2.271 87.123z" color="currentColor"/>`),
		g.Group(children),
	)
}