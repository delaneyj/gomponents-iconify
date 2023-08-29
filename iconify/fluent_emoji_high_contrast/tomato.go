package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tomato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.99 6a1 1 0 1 0-2 0v2.99h-.938A5.646 5.646 0 0 0 9.74 7H3.98a5.64 5.64 0 0 0 1.663 4.01A10.987 10.987 0 0 0 1 19.99c0 6.083 4.928 11 11 11h7.96c6.083 0 11-4.928 11-11a10.967 10.967 0 0 0-4.647-8.986A5.642 5.642 0 0 0 27.97 7h-5.76a5.646 5.646 0 0 0-4.312 1.99h-.908V6ZM6.054 13.232a5.666 5.666 0 0 0-.304 1.838h5.77a5.646 5.646 0 0 0 4.468-2.188a5.648 5.648 0 0 0 4.472 2.188h5.76c0-.642-.107-1.259-.305-1.835a8.957 8.957 0 0 1 3.045 6.753v.002c0 4.968-4.023 9-9 9H12c-4.968 0-9-4.023-9-9a8.98 8.98 0 0 1 3.054-6.758Z"/>`),
		g.Group(children),
	)
}