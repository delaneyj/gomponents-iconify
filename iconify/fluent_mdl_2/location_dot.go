package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 640q79 0 149 30t122 83t82 122t31 149q0 79-30 149t-83 122t-122 82t-149 31q-79 0-149-30t-122-83t-82-122t-31-149q0-79 30-149t83-122t122-82t149-31z"/>`),
		g.Group(children),
	)
}