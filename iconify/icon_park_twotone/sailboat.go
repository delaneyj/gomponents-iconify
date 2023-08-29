package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSailboat0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="m8 28l-2-4h36l-2 4"/><path fill="#555" d="M28 4L10 24h18V4Z"/><path d="m28 11l10 13H28V11ZM4 34l2.5 1.351c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0L44 34M4 40l2.5 1.351c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0L44 40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSailboat0)"/>`),
		g.Group(children),
	)
}