package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m247 245l10-85L122 23c29-14 62-23 97-23c121 0 219 97 219 219c0 18-3 38-7 56l249 248c45 46 49 116 8 157s-112 38-158-8L287 427c-22 7-44 12-68 12C98 439 0 340 0 219c0-34 7-64 20-93l139 141z"/>`),
		g.Group(children),
	)
}