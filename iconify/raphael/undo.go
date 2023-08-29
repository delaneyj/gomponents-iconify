package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.98 9.073V6.817L.876 13.807l12.106 6.99v-2.422c3.286-.002 9.053.28 9.053 2.27c0 2.78-6.023 4.262-6.023 4.262v2.132s13.53.462 13.53-9.824c0-8.082-11.588-8.385-16.56-8.143z"/>`),
		g.Group(children),
	)
}