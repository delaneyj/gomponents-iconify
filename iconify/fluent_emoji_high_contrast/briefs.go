package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.36 10A2.36 2.36 0 0 0 1 12.373v5.257v-.003v.001a5.48 5.48 0 0 0 1.639 3.925l7.936 7.826A5.53 5.53 0 0 0 14.442 31l.008-1v1h3.101a5.53 5.53 0 0 0 3.87-1.587l7.94-7.83A5.478 5.478 0 0 0 31 17.624v-5.252A2.36 2.36 0 0 0 28.64 10H3.36ZM1 12.37v.007v-.007Zm2 1.627l7.27.002l-3.637 8.684l-2.59-2.555l-.002-.001A3.479 3.479 0 0 1 3 17.633v-3.636Zm5.168 10.2L12.438 14l7.122.003l4.174 10.32l-3.716 3.664A3.53 3.53 0 0 1 17.549 29h-3.094a3.53 3.53 0 0 1-2.466-1.036l-3.82-3.768Zm17.108-1.395l-3.558-8.798l7.282.003v3.633a3.481 3.481 0 0 1-1.04 2.517l-.002.001l-2.681 2.644Z"/>`),
		g.Group(children),
	)
}