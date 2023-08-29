package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m18.621 2.55l2.829 2.83l-1.414 1.414l-2.829-2.828l1.414-1.415ZM12.822 8.6h-2v4h2v-4Z"/><path fill-rule="evenodd" d="M5.186 18.814A9 9 0 1 0 17.914 6.086A9 9 0 0 0 5.186 18.814Zm1.415-1.415A7 7 0 1 0 16.5 7.5a7 7 0 0 0-9.9 9.9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}