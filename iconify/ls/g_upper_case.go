package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M397.018 475v-73h397c-1 25-5 49-11 73c-36 145-157 265-319 294c-155 28-305-37-391-155c-33-45-57-99-67-158c-38-210 108-411 324-449c156-28 306 36 391 153l-63 44c-58-79-153-131-261-131c-178 0-322 142-322 315c0 69 22 131 60 183c58 79 154 130 262 130c147 0 272-95 311-226h-311z"/>`),
		g.Group(children),
	)
}