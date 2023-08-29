package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.473 22.94c.441 8.824-4.626 17.007-12.75 20.592c-8.125 3.583-17.63 1.829-23.918-4.414A21.274 21.274 0 0 1 4.341 15.35C7.939 7.272 16.167 2.227 25.045 2.656"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.344 34.573c-3.91-4.125-5.1-10.367-3.014-15.799c2.086-5.432 7.036-8.978 12.53-8.977c5.493 0 10.442 3.549 12.526 8.982c2.085 5.433.892 11.673-3.019 15.797"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.34 29.11c-1.885-1.927-2.459-4.842-1.453-7.38c1.006-2.537 3.392-4.194 6.04-4.194c2.648.001 5.033 1.658 6.038 4.196s.43 5.453-1.455 7.38"/>`),
		g.Group(children),
	)
}