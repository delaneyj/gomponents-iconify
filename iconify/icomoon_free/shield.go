package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 0L8 2L1 0S.93.808 1 2l7 2.189L15 2c.07-1.192 0-2 0-2zM1.128 3.049C1.503 6.966 2.901 13.553 8 16c5.099-2.448 6.497-9.034 6.872-12.951L8 5.633L1.128 3.049z"/>`),
		g.Group(children),
	)
}