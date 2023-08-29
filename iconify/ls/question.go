package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Question(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M86 281H12c-7-23-12-46-12-71C0 84 81 0 207 0c117 0 191 64 204 178c1 8 1 17 1 26c0 53-14 98-50 135c-35 35-59 58-97 91c-41 36-41 98-41 144h-77c0-108 24-153 59-188c34-34 83-64 109-97c22-26 28-57 28-85c0-87-49-136-136-136c-88 0-138 55-138 142c0 26 6 49 17 71zm61 373h77v77h-77v-77z"/>`),
		g.Group(children),
	)
}