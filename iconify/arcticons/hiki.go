package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hiki(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm2 17.706L16.224 42.5M5.5 33.675L16.224 42.5M5.71 14.963l29.18-1.532M38.85 5.5L22.775 37.69M34.89 13.431l7.61 9.775M24.893 36.135L42.5 38.088M15.047 14.472l6.43 24.171M5.5 28.755l9.547-14.283m-8.55-8.693l8.55 8.693"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.792 19.634c1.675-3.354-1.185-5.848-3.681-5.848s-5.41 1.597-5.41 3.616c-.841-1.99-1.706-3.242-3.844-3.13c-3.161.166-2.687 4.419-2.053 6.805"/>`),
		g.Group(children),
	)
}