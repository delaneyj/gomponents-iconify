package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLanguageZiTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M17.25 1a.75.75 0 0 1 .75.75V3h4.5a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0V4.5h-9.5v.75a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 11.5 3h5V1.75a.75.75 0 0 1 .75-.75zm-3.5 5a.75.75 0 0 0 0 1.5h4.428L16.77 8.674a.75.75 0 0 0-.27.576v.25h-4.75a.75.75 0 0 0 0 1.5h4.75v2.734a.75.75 0 0 1-.966.718l-.569-.17a.75.75 0 0 0-.43 1.436l.569.171A2.25 2.25 0 0 0 18 13.734V11h4.25a.75.75 0 0 0 0-1.5h-4.128l2.608-2.174A.75.75 0 0 0 20.25 6h-6.5zm-5.361.477l-.049-.104a.73.73 0 0 0-1.315.087l-5.964 14.5l-.032.096a.754.754 0 0 0 .426.886a.73.73 0 0 0 .963-.402l1.547-3.76l.094.006h7.087l1.433 3.737l.042.092a.73.73 0 0 0 .91.334a.755.755 0 0 0 .418-.972l-5.56-14.5zm-3.74 9.809L7.81 8.747l2.947 7.539h-6.11z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}