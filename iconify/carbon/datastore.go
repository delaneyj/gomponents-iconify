package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datastore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="23" cy="23" r="1" fill="currentColor"/><path fill="currentColor" d="M8 22h12v2H8z"/><circle cx="23" cy="9" r="1" fill="currentColor"/><path fill="currentColor" d="M8 8h12v2H8z"/><path fill="currentColor" d="M26 14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2v4H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-2v-4ZM6 6h20v6H6Zm20 20H6v-6h20Zm-4-8H10v-4h12Z"/>`),
		g.Group(children),
	)
}