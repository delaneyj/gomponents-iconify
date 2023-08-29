package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sealnote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.524 4.5h30.952v39h-23.11l-7.842-8.64V4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.365 43.5v-8.667H8.524"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="16.387" r="1.912"/><path d="M26.823 22.688a5.624 5.624 0 0 0 2.458-6.557a5.546 5.546 0 0 0-5.9-3.778a5.55 5.55 0 0 0-1.117 10.79m4.559-.455l.002 10.938l-1.768 1.97L22.83 33.6l1.12-1.762l-2.242-2.275l1.619-2.17l-1.975-1.838l.915-2.342"/></g>`),
		g.Group(children),
	)
}