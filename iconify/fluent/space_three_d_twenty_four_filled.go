package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceThreeDTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.25A3.25 3.25 0 0 1 6.25 3h11.5A3.25 3.25 0 0 1 21 6.25V15h-.813l-2.25-3H19.5V6.25a1.75 1.75 0 0 0-1.75-1.75H6.25A1.75 1.75 0 0 0 4.5 6.25V12h1.563l-2.25 3H3V6.25ZM15.692 16.5H21v1.25A3.25 3.25 0 0 1 17.75 21h-.667l-.367-1.224l-1.024-3.276Zm2.62-1.5h-3.089l-.937-3h1.777l2.25 3Zm-7.789 0h3.129l-.938-3h-1.44l-.75 3Zm3.598 1.5h-3.973L9.023 21h6.494l-.234-.78l-1.162-3.72ZM9.726 12l-.75 3h-3.29l2.25-3h1.79Zm-1.125 4.5H3v1.25A3.25 3.25 0 0 0 6.25 21h1.227l1.125-4.5Z"/>`),
		g.Group(children),
	)
}