package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparklingTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 1.208l1.32 2.473L20.792 5L18.32 6.319L17.002 8.79L15.68 6.32L13.208 5l2.473-1.318L17 1.208Zm-6.333 8.125l5 2.667l-5 2.666l-2.666 5l-2.667-5l-5-2.666l5-2.667l2.667-5l2.666 5Zm.75 2.667L9.19 10.81L8.001 8.583l-1.189 2.228L4.584 12l2.228 1.188l1.189 2.228l1.188-2.228L11.417 12Zm8.25 4.333l-1.666-3.125l-1.667 3.125L13.209 18l3.125 1.666l1.667 3.125l1.666-3.125L22.792 18l-3.125-1.667Z"/>`),
		g.Group(children),
	)
}