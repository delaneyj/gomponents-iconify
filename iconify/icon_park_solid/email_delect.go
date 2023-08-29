package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailDelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEmailDelect0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H4v30h20"/><path d="m4 9l20 15L44 9"/><path fill="#fff" d="M32 31h10l-2 10h-6l-2-10Z"/><path d="m36 31l2-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEmailDelect0)"/>`),
		g.Group(children),
	)
}