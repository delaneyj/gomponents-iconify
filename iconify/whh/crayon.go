package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crayon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m965.975 1020l-124-71q-18 12-39.5 10t-37.5-18l-746-746q-18-18-18-44t18-44l89-88q18-19 44-19t44 19l746 746q15 15 17 36.5t-9 39.5l71 124q11 33-5.5 49.5t-49.5 5.5zm-689-826l-83 82l44 43l82-82zm96 96l-83 82l235 235l82-83zm287 287l-82 82l43 44l82-83z"/>`),
		g.Group(children),
	)
}