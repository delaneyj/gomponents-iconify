package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Five(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m203 72l-36 192c20-6 42-9 65-9c137 0 248 112 248 249S369 752 232 752C115 752 27 672 0 562h73c25 71 81 122 160 122c99 0 179-80 179-179c0-98-80-180-179-180c-26 0-53 6-81 14c-30 8-51 14-78 27L144 0h317v72H203z"/>`),
		g.Group(children),
	)
}