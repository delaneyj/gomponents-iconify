package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21.786 12.876l7.556-4.363l-7.556-4.363v2.598H2.813v3.5h18.973v2.628zm-11.418 5.248l-7.556 4.362l7.556 4.362V24.25h18.974v-3.5H10.368v-2.626z"/>`),
		g.Group(children),
	)
}