package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLowest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m3 14l.53-.53l2.25-2.25a.75.75 0 0 0-1.06-1.061l-.97.97v-8.38a.75.75 0 0 0-1.5 0v8.38l-.97-.97a.75.75 0 1 0-1.06 1.06l2.25 2.25L3 14Zm5-9.25c0 .414.336.75.75.75h1.5a.75.75 0 0 0 0-1.5h-1.5a.75.75 0 0 0-.75.75Zm.75 4a.75.75 0 1 1 0-1.5h4.5a.75.75 0 0 1 0 1.5h-4.5Zm0 3.25a.75.75 0 0 1 0-1.5h6.5a.75.75 0 0 1 0 1.5h-6.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}