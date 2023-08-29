package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5.07A7.005 7.005 0 0 0 5.07 11H7v2H5.07A7.004 7.004 0 0 0 11 18.93V17h2v1.93A7.004 7.004 0 0 0 18.93 13H17v-2h1.93A7.004 7.004 0 0 0 13 5.07V7h-2V5.07ZM3.055 11A9.004 9.004 0 0 1 11 3.055V1h2v2.055A9.004 9.004 0 0 1 20.945 11H23v2h-2.055A9.004 9.004 0 0 1 13 20.945V23h-2v-2.055A9.004 9.004 0 0 1 3.055 13H1v-2h2.055ZM15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`),
		g.Group(children),
	)
}