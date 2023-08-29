package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogTwentyEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM29 35.5H11m26 6.954v-4.74h-2.234c-1.117 0-2.016-.9-2.016-2.016v-.538c0-1.117.9-2.016 2.016-2.016H37v-6.68h-2.484c-1.117 0-2.016-.9-2.016-2.016v-.538c0-1.117.9-2.016 2.016-2.016H37v-6.93h-2.484c-1.117 0-2.016-.9-2.016-2.016v-.538c0-1.116.9-2.015 2.016-2.015H37V5.54M29 24.5H11m13-11H11"/>`),
		g.Group(children),
	)
}