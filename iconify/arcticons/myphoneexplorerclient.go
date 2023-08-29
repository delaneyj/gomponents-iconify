package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myphoneexplorerclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.76 20.46c1 11.61-2.44 15.9-18.33 14.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.56 42.9l-7.12-7.5l7.12-7.51m-18.32-.35c-1-11.61 2.44-15.9 18.33-14.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.47 5.13l7.12 7.5l-7.12 7.51"/>`),
		g.Group(children),
	)
}