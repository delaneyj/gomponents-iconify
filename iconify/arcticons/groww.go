package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Groww(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.37 36.31l12.1-6.8l12.94 5.3l13.86-13.92"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.23 29.56L18.47 21l12.94 5.3l11.84-11.89M12.13 41.93l7.12-4.01l12.94 5.3l12.13-12.17"/>`),
		g.Group(children),
	)
}