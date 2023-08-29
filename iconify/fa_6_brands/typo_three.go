package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypoThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M178.7 78.4c0-24.7 5.4-32.4 13.9-39.4c-69.5 8.5-149.3 34-176.3 66.4c-5.4 7.7-9.3 20.8-9.3 37.1C7 246 113.8 480 191.1 480c36.3 0 97.3-59.5 146.7-139c-7 2.3-11.6 2.3-18.5 2.3c-57.2 0-140.6-198.5-140.6-264.9zM301.5 32c-30.1 0-41.7 5.4-41.7 36.3c0 66.4 53.8 198.5 101.7 198.5c26.3 0 78.8-99.7 78.8-182.3c0-40.9-67-52.5-138.8-52.5z"/>`),
		g.Group(children),
	)
}