package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Successfactors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.349 27.29c-11.17-10.956 4.397-28.4 15.878-15.668c10.76-12.732 26.58 4.543 15.505 15.668L24.227 42.865L8.35 27.29Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.267 28.468a7.782 7.782 0 1 1 10.004 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.939 28.032a6.485 6.485 0 0 1 7.888-10.28m11.444 10.716c.847.311 1.722.786 2.658 1.638m-12.662-1.638c-2.93 2.01-5.207 4.596-6.613 7.95m-2.715-8.386c-1.52.907-2.427 1.815-2.962 2.722"/>`),
		g.Group(children),
	)
}