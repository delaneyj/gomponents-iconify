package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Settingsalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534zm8.386 13.502a5.36 5.36 0 0 1-5.685 1.586l-7.187 8.266a2.113 2.113 0 0 1-3.187-2.775l7.198-8.274a5.348 5.348 0 0 1 6.137-7.497l-2.755 3.212l.9 2.62l2.723.53l2.76-3.22a5.339 5.339 0 0 1-.902 5.553z"/>`),
		g.Group(children),
	)
}