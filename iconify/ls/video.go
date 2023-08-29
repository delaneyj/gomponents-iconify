package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m496 258l219-126c29-17 53-5 53 30v372c0 34-24 48-53 31L496 438v94c0 35-28 63-63 63H61c-34 0-61-28-61-63V161c0-35 27-62 61-62h372c35 0 63 27 63 62v97z"/>`),
		g.Group(children),
	)
}