package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetStraightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 42h-40a14 14 0 0 0-14 14v88a18 18 0 0 1-36 0V56a14 14 0 0 0-14-14H56a14 14 0 0 0-14 14v88a86 86 0 0 0 86 86h.65c47.06-.35 85.35-39.38 85.35-87V56a14 14 0 0 0-14-14Zm-40 12h40a2 2 0 0 1 2 2v34h-44V56a2 2 0 0 1 2-2ZM56 54h40a2 2 0 0 1 2 2v34H54V56a2 2 0 0 1 2-2Zm72.56 164H128a74 74 0 0 1-74-74v-42h44v42a30 30 0 0 0 60 0v-42h44v41c0 41.05-32.94 74.7-73.44 75Z"/>`),
		g.Group(children),
	)
}