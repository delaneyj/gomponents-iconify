package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeftRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.5 6.002a1.5 1.5 0 0 0-.95-1.396a.75.75 0 1 1 .493-1.413v-.005A3.001 3.001 0 0 1 17 6.002v8c0 1.29-.814 2.39-1.957 2.814v-.004a.75.75 0 0 1-.586-1.38l.01-.004a1.5 1.5 0 0 0 1.033-1.426v-8Zm-6.749-3a.75.75 0 1 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Zm0 12.5a.75.75 0 1 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5ZM5.449 4.606A1.5 1.5 0 0 0 4.5 6.002v8a1.5 1.5 0 0 0 1.033 1.426l.01.003a.75.75 0 0 1-.586 1.381v.004A3.001 3.001 0 0 1 3 14.002v-8c0-1.29.814-2.39 1.957-2.814v.005a.75.75 0 1 1 .492 1.413Z"/>`),
		g.Group(children),
	)
}