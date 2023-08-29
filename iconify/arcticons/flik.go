package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.776 19.541v9.449m4.724-.014l-3.362-3.348l2.915-2.901M11.318 28.99v-7.795c0-.913.74-1.653 1.653-1.653h0c.814 0 1.322.24 1.67.691M8.5 22.73h6.142M8.5 25.13h6.142"/><circle cx="28.871" cy="19.837" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.509 22.73h2.362v6.26m2.362 0h-4.724m-3.544 0h-4.724m2.362 0v-9.449h-2.362"/>`),
		g.Group(children),
	)
}