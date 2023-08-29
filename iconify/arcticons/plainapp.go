package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plainapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 29.6a7.4 7.4 0 0 1 14.8-.006v.006a7.4 7.4 0 0 1-14.8.002V29.6m20.3 8.3a4.65 4.65 0 1 1 9.3 0a4.65 4.65 0 0 1-9.3 0m-5.5-21.3c0-6.13 4.97-11.1 11.1-11.1s11.1 4.97 11.1 11.1s-4.97 11.1-11.1 11.1s-11.1-4.97-11.1-11.1"/>`),
		g.Group(children),
	)
}