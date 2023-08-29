package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Applicationsinfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.45 8.4a2 2 0 0 0-1.95 1.95v27.3a2 2 0 0 0 2 2h35.1a2 2 0 0 0 2-2V21.9a1.93 1.93 0 0 0-.56-1.38L31.41 9A2 2 0 0 0 30 8.4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.12 12.79A11.21 11.21 0 1 1 7.91 24a11.21 11.21 0 0 1 11.21-11.21Z"/><circle cx="19.12" cy="18.22" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.12 22.08v8.45"/>`),
		g.Group(children),
	)
}