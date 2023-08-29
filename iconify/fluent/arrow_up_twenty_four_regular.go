package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.209 10.733a.75.75 0 0 0 1.086 1.034l5.954-6.251V20.25a.75.75 0 0 0 1.5 0V5.516l5.955 6.251a.75.75 0 0 0 1.087-1.034l-7.067-7.42a.995.995 0 0 0-.58-.3a.754.754 0 0 0-.29.001a.995.995 0 0 0-.578.3l-7.067 7.419Z"/>`),
		g.Group(children),
	)
}