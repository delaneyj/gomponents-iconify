package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.421 0L5 15.127h2.138L10.56 0H8.421Zm8.993 8.873l-1.496 6.225h2.138l1.496-6.225h-2.138Zm-3.635 0L10.36 24h2.138l3.42-15.127H13.78Zm-3.647 0l-1.497 6.225h2.138l1.496-6.225h-2.137Z"/>`),
		g.Group(children),
	)
}