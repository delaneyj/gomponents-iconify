package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UfoThreeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M18.922 12.15c1.603 2.623 2.072 5 .98 6.091c-1.63 1.63-6.118-.214-10.023-4.12C5.974 10.217 4.129 5.73 5.759 4.1c1.092-1.092 3.468-.624 6.092.98"/><path stroke="currentColor" stroke-width="1.5" d="M11.485 5.445a4.729 4.729 0 0 1 6.687 0l.384.383a4.729 4.729 0 0 1 0 6.688c-.126.125-.286.21-.46.178c-.497-.09-1.72-.56-3.975-2.815c-2.255-2.256-2.725-3.477-2.814-3.974c-.032-.174.052-.335.178-.46Z" opacity=".5"/><circle cx="12" cy="12" r=".5" fill="currentColor" stroke="currentColor" opacity=".5" transform="rotate(45 12 12)"/><circle cx="9.172" cy="7.757" r="1" fill="currentColor" transform="rotate(45 9.172 7.757)"/><circle cx="16.243" cy="14.828" r="1" fill="currentColor" transform="rotate(45 16.243 14.828)"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m3.636 20.363l4.121-4.12"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m10 22l2.707-2.222M2 14l2.222-2.707" opacity=".5"/></g>`),
		g.Group(children),
	)
}