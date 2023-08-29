package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m159 345l172 173l-83 80L0 350L254 96l79 78zm450 4L437 176l83-80l248 248l-254 254l-79-78z"/>`),
		g.Group(children),
	)
}