package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArchiveTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.519 10a6.5 6.5 0 0 1 12.962 0h.269a4.751 4.751 0 0 1 4.691 4H15a2 2 0 0 0-2 2v1a2 2 0 0 0 1 1.732v.768H7.25a4.75 4.75 0 1 1 0-9.5h.269ZM14 16a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H15a1 1 0 0 1-1-1v-1Zm12 3H15v5a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3v-5Zm-7.5 2h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}