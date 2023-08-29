package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5zM320 184v79L118 61q17-8 31-12V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5zm-6 181H0v-21l43-43V184q0-38 19-71L0 51l27-27l357 357l-27 27z"/>`),
		g.Group(children),
	)
}