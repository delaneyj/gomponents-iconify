package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="m579 141l-53 509q-2 21-21 39t-51 31t-74 22t-90 8t-90-8t-74-22t-51-31t-22-39L0 141h1q-1-2-1-5q0-27 23-50t62-41t92-27T290 8t113 10t92 27t62 41t22 50v5zm-289 76q43 0 81-7t66-17t45-26t16-31t-16-32t-45-25l-66-18q-36-10-81-6q-44 0-82 6t-66 18t-44 25t-17 32t17 31t44 26t66 17t82 7z"/>`),
		g.Group(children),
	)
}