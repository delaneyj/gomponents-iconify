package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetTopStackTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M9.5 5.85a.75.75 0 0 1 .67.415l.821 1.64l1.602.2a.75.75 0 0 1 .387 1.321l-1.149.957l.401 1.804a.75.75 0 0 1-1.13.799L9.5 11.984l-1.603 1.002a.75.75 0 0 1-1.13-.799l.402-1.803l-1.15-.958a.75.75 0 0 1 .388-1.32l1.602-.2l.82-1.641A.75.75 0 0 1 9.5 5.85zm0 2.427l-.33.658a.75.75 0 0 1-.577.41l-.286.035l.173.144a.75.75 0 0 1 .252.739l-.112.502l.482-.301a.75.75 0 0 1 .796 0l.482.301l-.112-.502a.75.75 0 0 1 .252-.74l.173-.143l-.286-.036a.75.75 0 0 1-.578-.409L9.5 8.277z" fill="currentColor"/><path d="M4 4h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2zm11 1.5H4a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5z" fill="currentColor"/><path d="M4.563 16A2 2 0 0 0 6.5 17.5h9a4 4 0 0 0 4-4v-5A2 2 0 0 0 18 6.563V13.5a2.5 2.5 0 0 1-2.5 2.5H4.563z" fill="currentColor"/><path d="M7.063 18.5A2 2 0 0 0 9 20h7.25A5.75 5.75 0 0 0 22 14.25V11a2 2 0 0 0-1.5-1.937v5.187a4.25 4.25 0 0 1-4.25 4.25H7.063z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}