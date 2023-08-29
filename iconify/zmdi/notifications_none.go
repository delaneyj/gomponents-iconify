package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5zM320 301l43 43v21H0v-21l43-43V184q0-49 30-86.5T149 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5v117zm-43 22V184q0-40-28-68t-68-28t-68 28t-28 68v139h192z"/>`),
		g.Group(children),
	)
}