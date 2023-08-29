package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SailboatOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSailboatOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M21 31V5L7 31h14Zm6 0V13l14 18H27Z"/><path d="M5 37h38l-5 6H10l-5-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSailboatOne0)"/>`),
		g.Group(children),
	)
}