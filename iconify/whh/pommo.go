package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pommo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 768V354L608 768H416L192 354v414H0V0h192l320 576L832 0h192v768H832z"/>`),
		g.Group(children),
	)
}