package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldDismissTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.45 2.15C14.992 4.057 17.587 5 20.25 5a.75.75 0 0 1 .75.75V11c0 5.001-2.958 8.676-8.725 10.948a.75.75 0 0 1-.55 0C5.958 19.676 3 16 3 11V5.75A.75.75 0 0 1 3.75 5c2.663 0 5.258-.943 7.8-2.85a.75.75 0 0 1 .9 0ZM9.28 8.222a.75.75 0 0 0-1.13.975l.072.084l2.723 2.723l-2.723 2.725a.749.749 0 1 0 1.059 1.059l2.723-2.724l2.725 2.724a.75.75 0 0 0 .975.073l.084-.073a.75.75 0 0 0 .073-.975l-.073-.084l-2.724-2.725l2.724-2.723a.749.749 0 1 0-1.06-1.06l-2.724 2.724l-2.723-2.723Z"/>`),
		g.Group(children),
	)
}