package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoHearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M27.12 31.749c1.277-7.515-10.99-19.654-21.333-9.734C-5.737 33.062 12.649 57.156 13.42 60c2.5-.722 32.946-1.499 34.413-17.237C49.194 28.14 31.497 26.258 27.12 31.749M59.198 5.608C52.55.555 46.343 8.681 47.677 12.477c-3.242-2.877-12.77-.492-11.088 7.568c1.877 8.969 19.01 7.153 20.611 7.623c.403-1.442 9.151-16.635 1.998-22.06"/>`),
		g.Group(children),
	)
}