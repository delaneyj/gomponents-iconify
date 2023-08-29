package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.29 47.91A20 20 0 0 0 216 36H40a20 20 0 0 0-14.8 33.45l.12.14L92 140.75V216a20 20 0 0 0 31.1 16.64l32-21.33a20 20 0 0 0 8.9-16.65v-53.91l66.67-71.16l.12-.14a20 20 0 0 0 3.5-21.54Zm-88.88 77.58a19.93 19.93 0 0 0-5.41 13.68v53.35l-24 16v-69.35a19.93 19.93 0 0 0-5.41-13.68L49.23 60h157.54Z"/>`),
		g.Group(children),
	)
}