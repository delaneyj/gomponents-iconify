package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="27.97" cy="16.163" r="1.663"/><path d="M27.97 24v4.75a4.75 4.75 0 0 0 4.75 4.75M15.28 24H24m-8.72 9.5V19.25a4.75 4.75 0 0 1 4.75-4.75H24"/></g>`),
		g.Group(children),
	)
}