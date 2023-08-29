package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BritishPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17.688 5c-.27 0-.548.028-.813.063c-3.177.41-5.875 3.14-5.875 6.592c0 1.13.283 2.242.656 3.344H9v2h3.406c.198.554.36 1.102.5 1.655c.563 2.206.62 4.182-1.375 6.344H8v2h16v-5h-2v3h-7.938c1.393-2.307 1.318-4.746.782-6.845c-.102-.4-.225-.772-.344-1.156H20v-2h-6.156C13.39 13.692 13 12.53 13 11.655c0-3.232 3.3-5.543 6.375-4.312l.75-1.844c-.815-.326-1.63-.498-2.438-.5z"/>`),
		g.Group(children),
	)
}