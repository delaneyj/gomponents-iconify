package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThongSandal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.5 2a7.502 7.502 0 0 1 7.073 5C24.145 13.43 24 23 24 23a7 7 0 1 1-14 0c0-5.653-1.304-8.538-1.806-9.44A7.461 7.461 0 0 1 7 9.5A7.5 7.5 0 0 1 14.5 2Zm.94 6.474l-.384-.119a.845.845 0 0 1-.534-.489L14 6.584a.929.929 0 0 0-1.724.69l.521 1.281a.937.937 0 0 1-.126.926l-.801 1.034a8.158 8.158 0 0 0-1.614 7.141l1.003 4.012a1 1 0 0 0 1.94-.485l-1.003-4.012a6.152 6.152 0 0 1 1.216-5.383l.84-1.02a.889.889 0 0 1 .975-.248a7.92 7.92 0 0 1 5.105 7.42v2.457c0 .538.414 1.002.952 1.027a1 1 0 0 0 1.048-.999v-2.486a9.877 9.877 0 0 0-1.906-5.854a9.882 9.882 0 0 0-4.987-3.611Z"/>`),
		g.Group(children),
	)
}