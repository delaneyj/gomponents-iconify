package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 12L5.002 6.335c-.43-.347-1.063-.283-1.366.178a9.99 9.99 0 0 0-1.295 2.898C.91 14.746 4.077 20.23 9.41 21.66c5.335 1.43 10.819-1.737 12.248-7.071C23.09 9.253 19.923 3.77 14.59 2.34a9.968 9.968 0 0 0-7.502.95"/><path fill="currentColor" d="M9.866 5.595a.75.75 0 1 0 .474 1.423l-.474-1.423ZM17.25 12c0 2.9-2.35 5.25-5.25 5.25v1.5A6.75 6.75 0 0 0 18.75 12h-1.5ZM12 17.25A5.25 5.25 0 0 1 6.75 12h-1.5A6.75 6.75 0 0 0 12 18.75v-1.5Zm0-10.5c2.9 0 5.25 2.35 5.25 5.25h1.5A6.75 6.75 0 0 0 12 5.25v1.5ZM6.75 12c0-1.269.45-2.43 1.198-3.338L6.79 7.707A6.725 6.725 0 0 0 5.25 12h1.5Zm3.59-4.982A5.244 5.244 0 0 1 12 6.75v-1.5c-.745 0-1.462.12-2.134.345l.474 1.423Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}