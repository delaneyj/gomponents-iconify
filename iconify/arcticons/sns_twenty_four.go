package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnsTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.025 44.658a15.014 15.014 0 0 1-2.768-8.882c0-4.795 3.995-16.36 20.145-16.36a12.598 12.598 0 0 1 10.096 4.856M2.741 27.227a3.72 3.72 0 0 0 3.222 2.156c2.733 0 2.954-2.796 4.794-7.191c5.39-12.878 16.36-14.678 22.794-14.678a18.022 18.022 0 0 1 5.177.823"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.986 31.632c2.777 3.053 6.59 4.606 11.972 4.606c4.954 0 11.624-2.774 16.324-5.087M15.257 4.353c-1.979 3.997-4.38 10.768-3.953 16.609"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.499 19.471c3.166 4.954 2.938 11.012 2.075 15.305M12.027 13.362a30.015 30.015 0 0 1 6.41-1.325"/>`),
		g.Group(children),
	)
}