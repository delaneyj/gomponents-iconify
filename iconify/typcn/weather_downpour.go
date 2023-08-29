package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherDownpour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1zm-6 0a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1zm3 2a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1zm-6-6c-2.206 0-4-1.794-4-4a4.007 4.007 0 0 1 3.001-3.874L5 10c0-3.309 2.691-6 6-6a5.97 5.97 0 0 1 5.65 4.015C19.586 7.771 22 10.128 22 13a5.011 5.011 0 0 1-3.666 4.819a1 1 0 1 1-.532-1.927A3.008 3.008 0 0 0 20 13c0-1.654-1.346-3-3-3c-.242 0-.499.041-.811.13l-1.074.306l-.185-1.102A3.98 3.98 0 0 0 11 6a4.004 4.004 0 0 0-3.918 4.808l.248 1.202l-1.422-.016C4.897 12 4 12.897 4 14s.897 2 2 2a1 1 0 1 1 0 2z"/>`),
		g.Group(children),
	)
}