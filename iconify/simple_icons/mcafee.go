package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mcafee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4.823L1.58 0v19.177L12 24l10.42-4.823V0zm6.172 11.626l-6.143 2.843l-6.144-2.843V6.69l6.144 2.842l6.143-2.842z"/>`),
		g.Group(children),
	)
}