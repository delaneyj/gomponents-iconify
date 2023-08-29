package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolidLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M461.596 17.701C271.149 128.454-109.411-4.192 66.276 350.885c3.065 5.77 6.619 11.164 10.529 16.293c-29.888 33.096-51.12 70.802-57.117 114.554c5.26 3.375 14.588 7.464 26.88 9.916c13.06 2.605 29.481 3.516 47.916 1.711c-11.483-36.045-7.774-70.234 5.836-101.043c5.5 6.543 10.283 10.23 14.782 13.012C581.497 693.816 499.604 120.792 461.596 17.701z"/>`),
		g.Group(children),
	)
}