package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indoormap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.872 18.607L24 6.436L7.128 18.607v22.957h33.744V18.607zM4.5 20.503l2.628-1.896M43.5 20.503l-2.628-1.896"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.929 30.479h6.141v11.085h-6.141zm10.053 0h6.141v6.141h-6.141zm3.267-16.65V9.104h-6.551M10.877 30.479h6.141v6.141h-6.141z"/>`),
		g.Group(children),
	)
}