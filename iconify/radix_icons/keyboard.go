package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 4h-12a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5Zm-12-1A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12h12a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 13.5 3h-12ZM2 5h1v1H2V5Zm3 0H4v1h1V5Zm1 0h1v1H6V5Zm3 0H8v1h1V5Zm1 0h1v1h-1V5Zm3 0h-1v1h1V5Zm-2 2h1v1h-1V7Zm2 2h-1v1h1V9ZM9 7h1v1H9V7ZM8 7H7v1h1V7ZM5 7h1v1H5V7ZM4 7H3v1h1V7ZM2 9h1v1H2V9Zm9 0H4v1h7V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}