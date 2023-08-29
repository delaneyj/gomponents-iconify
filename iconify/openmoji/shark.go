package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M18.25 40.965S15 46.828 11 49.931c0 0 2-4.138 1-8.276s-8-9.31-8-9.31s10 0 14 5.172c0 0 17 1.035 21-6.207c0 0 1-5.172-3-9.31c0 0 5 0 10 4.138c4.846 4.01 3.116 4.134 9.38 7.905a6.13 6.13 0 0 0 1.3.588L69 38.55s0 4.139-14 6.208c0 0-1 5.172-5 7.241c0 0 2-6.207-1-6.207c0 0-2 0-9-1.034"/><path fill="#D0CFCE" d="M37.505 42.314s-3.132 1.905-3.679 4.455c0 0-2.866-3.02-4.972-2.56l4.759-3.47M44.556 37s4.25 2 0 5"/><path fill="#D0CFCE" d="M46.556 37s4.25 2 0 5"/><path fill="#D0CFCE" d="M48.556 37s4.25 2 0 5"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10"><path stroke-width="2" d="M40 44c7 1 9 1 9 1c3 0 1 6 1 6c4-2 5-7 5-7c14-2 14-6 14-6l-12.32-3.79a6.239 6.239 0 0 1-1.3-.568C49.115 29.996 50.845 29.877 46 26c-5-4-10-4-10-4c4 4 3 9 3 9c-4 7-21 6-21 6c-4-5-14-5-14-5s7 5 8 9s-1 8-1 8c4-3 7.25-8.667 7.25-8.667L27 42"/><path stroke-width="2" d="M37 43s-2.928 1.668-3.478 3.836c0 0-2.594-2.504-4.548-2.078l4.458-3.02"/><path d="M45 37s2 2 0 5m2-5s2 2 0 5m2-5s2 2 0 5"/></g>`),
		g.Group(children),
	)
}