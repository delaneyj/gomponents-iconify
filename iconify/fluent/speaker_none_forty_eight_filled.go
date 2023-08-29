package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNoneFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M25.081 6.418C26.208 5.408 28 6.207 28 7.721v32.558c0 1.514-1.792 2.314-2.919 1.302l-8.206-7.366A4.75 4.75 0 0 0 13.702 33H9a5.25 5.25 0 0 1-5.25-5.25v-7.5C3.75 17.35 6.1 15 9 15h4.702a4.75 4.75 0 0 0 3.173-1.215l8.206-7.367z" fill="currentColor"/><path d="M32.366 18.366a1.25 1.25 0 0 1 1.768 0L38 22.232l3.866-3.866a1.25 1.25 0 0 1 1.768 1.768L39.768 24l3.866 3.866a1.25 1.25 0 0 1-1.768 1.768L38 25.768l-3.866 3.866a1.25 1.25 0 1 1-1.768-1.768L36.232 24l-3.866-3.866a1.25 1.25 0 0 1 0-1.768z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}