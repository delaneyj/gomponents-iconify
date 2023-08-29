package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M227 507L97 594s27-72 29-124C49 429 0 362 0 287C0 162 139 60 310 60c142 0 262 71 299 166c-14-2-28-3-42-3h-14c-70 3-135 26-185 65c-55 44-85 106-83 168c1 19 5 40 12 58c-24 0-47-3-70-7zm481 127l-85-46c-17 4-37 7-55 8h-12c-109 0-197-63-200-143c-4-83 86-154 200-158h11c109 0 198 61 201 142c2 49-28 93-77 123c2 35 17 74 17 74z"/>`),
		g.Group(children),
	)
}