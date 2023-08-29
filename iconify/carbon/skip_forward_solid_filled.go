package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardSolidFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm2.486 14.874l-9 5A1 1 0 0 1 8 21V11a1 1 0 0 1 1.486-.874l9 5a1 1 0 0 1 0 1.748ZM24 22h-2V10h2Z"/><path fill="none" d="M22 10h2v12h-2zM8.493 21.862A1 1 0 0 1 8 21V11a1 1 0 0 1 1.486-.874l9 5a1 1 0 0 1 0 1.748l-9 5a1 1 0 0 1-.993-.012z"/>`),
		g.Group(children),
	)
}