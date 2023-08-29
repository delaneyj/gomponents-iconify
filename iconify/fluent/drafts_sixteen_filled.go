package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DraftsSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.98 1.998a2.621 2.621 0 0 0-3.71.003l-.78.783l3.706 3.707l.786-.788a2.621 2.621 0 0 0-.003-3.705ZM2.802 9.49l5.98-5.998l3.707 3.707l-5.977 5.995a1.5 1.5 0 0 1-.558.354l-3.969 1.416a.75.75 0 0 1-.958-.96l1.426-3.963a1.5 1.5 0 0 1 .349-.551ZM1.5 2h7.357L7.86 3H1.5a.5.5 0 0 1 0-1Zm0 2h5.363l-.997 1H1.5a.5.5 0 0 1 0-1Zm0 2h3.37l-.998 1H1.5a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}