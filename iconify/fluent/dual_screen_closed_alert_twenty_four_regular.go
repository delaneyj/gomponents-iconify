package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenClosedAlertTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.499 2h.002A4.01 4.01 0 0 1 20.5 5.999v2.52l1.384 1.66A.5.5 0 0 1 21.5 11h-10a.5.5 0 0 1-.384-.82L12.5 8.519v-2.52A4.01 4.01 0 0 1 16.499 2ZM5.75 5h5.853a5.016 5.016 0 0 0-.103.996V6.5h-5v13h8.25a.75.75 0 0 0 .75-.75v-4.92a2.997 2.997 0 0 0 1.5.129v4.791A2.25 2.25 0 0 1 14.75 21h-9a.75.75 0 0 1-.75-.75V5.75A.75.75 0 0 1 5.75 5Zm9.754 7.732a1.984 1.984 0 0 1-.727-.732h3.446a1.984 1.984 0 0 1-2.719.732Z"/>`),
		g.Group(children),
	)
}