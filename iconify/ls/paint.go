package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M326 337h82V28h-82v309zm131-87l85 85c14 14 14 36 0 50L275 651c-13 14-36 14-49 0L10 435c-13-13-13-36 0-49l266-267v259h181V250zm74 24s67 27 67 159c0 82-9 95-9 95s128 14 128-111c0-126-186-143-186-143z"/>`),
		g.Group(children),
	)
}