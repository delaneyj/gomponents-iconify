package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallapop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.283 25.17c-9.014-2.363-9.314-9.972-6.079-14.203c3.474-4.542 9.483-2.107 11.624-.4c1.87-6.213 11.089-9.286 18.704-1.136c7.616 8.15-2.805 16.634-2.805 16.634"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.861 18.718c9.61 4.143 2.882 16.498-3.33 15.496c.143 4.072-1.604 9.293-9.153 9.286c-4.003-.004-10.02-4.61-5.811-14.029"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.664 23.703C2.129 27 4.606 42.885 19.5 38.048"/>`),
		g.Group(children),
	)
}