package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.248 28.461h0a3.32 3.32 0 0 1-3.251-3.367v-2.188a3.32 3.32 0 0 1 3.251-3.367h0a3.32 3.32 0 0 1 3.252 3.367v2.188a3.32 3.32 0 0 1-3.252 3.367Zm-7.834-5.555a3.32 3.32 0 0 1 3.252-3.367h0m-3.252 0v8.922m-11.899-1.683a3.132 3.132 0 0 1-2.763 1.683h0A3.32 3.32 0 0 1 9.5 25.094v-2.188a3.32 3.32 0 0 1 3.252-3.367h0a3.32 3.32 0 0 1 3.251 3.367v1.178H9.5m15.012 2.694a3.132 3.132 0 0 1-2.764 1.683h0a3.32 3.32 0 0 1-3.252-3.367v-2.188a3.32 3.32 0 0 1 3.252-3.367h0a3.32 3.32 0 0 1 3.251 3.367v1.178h-6.503"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}