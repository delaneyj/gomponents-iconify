package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mewe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.06 28.974a2.07 2.07 0 0 1-1.818 1.07a2.146 2.146 0 0 1-2.14-2.14v-1.39a2.14 2.14 0 1 1 4.28 0v.749h-4.28m20.078 1.711a2.07 2.07 0 0 1-1.82 1.07a2.146 2.146 0 0 1-2.139-2.14v-1.39a2.14 2.14 0 1 1 4.279 0v.749h-4.279m-1.006-5.777l-2.139 8.557l-2.139-8.557l-2.14 8.557l-2.139-8.557"/><circle cx="29.937" cy="18.812" r=".75" fill="currentColor"/><circle cx="25.508" cy="18.812" r=".75" fill="currentColor"/><circle cx="34.366" cy="18.812" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 30.043v-8.557l4.279 8.557l4.278-8.557v8.557"/><circle cx="12.779" cy="18.812" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}