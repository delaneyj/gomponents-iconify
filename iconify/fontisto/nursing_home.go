package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NursingHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 8.853l2.886 10.115c2.738-.403 5.899-.633 9.113-.633s6.375.23 9.467.675l-.353-.042l2.886-10.115C14.497 4.628 8.858 4.405 0 8.853zm14.918 4.276h-2.071V15.2h-1.686v-2.071H9.09v-1.686h2.071V9.371h1.686v2.072h2.071z"/>`),
		g.Group(children),
	)
}