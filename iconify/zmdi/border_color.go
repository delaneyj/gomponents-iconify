package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M379 149L165 363H85v-80L299 69zm63-63l-42 42l-80-80l42-42q6-6 15-6t15 6l50 50q6 6 6 15t-6 15zM0 427h512v85H0v-85z"/>`),
		g.Group(children),
	)
}