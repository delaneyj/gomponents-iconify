package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagLockAccentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.75 2.5a1.752 1.752 0 0 1 1.75 1.75v5.462a2.73 2.73 0 0 1-.8 1.944l-.334.335A3.5 3.5 0 0 0 14 14v.049a2.5 2.5 0 0 0-2 2.45v3.837a2.753 2.753 0 0 1-3.69-.181l-4.468-4.451a2.749 2.749 0 0 1 0-3.888l8.5-8.51a2.727 2.727 0 0 1 1.943-.807h5.465Z"/>`),
		g.Group(children),
	)
}