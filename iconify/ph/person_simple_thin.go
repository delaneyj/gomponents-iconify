package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M227.43 130.06a4 4 0 0 1-5.49 1.37c-.38-.23-37.87-22.29-89.94-23.38v42.43l63 70.86a4 4 0 0 1-6 5.32L128 158l-61 68.66a4 4 0 0 1-6-5.32l63-70.86v-42.43c-52.07 1.09-89.56 23.15-89.94 23.38a4 4 0 0 1-4.12-6.86c1.67-1 41.6-24.57 98.06-24.57s96.39 23.57 98.06 24.57a4 4 0 0 1 1.37 5.49ZM100 48a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm8 0a20 20 0 1 0 20-20a20 20 0 0 0-20 20Z"/>`),
		g.Group(children),
	)
}