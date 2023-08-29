package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 3.4c0-.635.515-1.15 1.15-1.15h16.2c.635 0 1.15.515 1.15 1.15v11.472c0 .305-.121.597-.337.813l-2.722 2.722a1.15 1.15 0 0 1-.813.337h-4.54l-3.45 3.451c-.6.599-1.622.175-1.622-.672v-2.78H3.9a1.15 1.15 0 0 1-1.15-1.15V3.4Zm1.5.35v13.494h3.866c.635 0 1.15.515 1.15 1.15v1.802l2.614-2.615a1.15 1.15 0 0 1 .814-.337h4.539l2.517-2.517V3.75H4.25Zm6.219 3.2a.75.75 0 0 1 .75.75v4.272a.75.75 0 1 1-1.5 0V7.7a.75.75 0 0 1 .75-.75Zm5.016 0a.75.75 0 0 1 .75.75v4.272a.75.75 0 1 1-1.5 0V7.7a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}