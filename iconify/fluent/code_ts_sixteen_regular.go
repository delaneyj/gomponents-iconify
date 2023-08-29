package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTsSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 9.75A2.28 2.28 0 0 1 10.5 12A2.28 2.28 0 0 1 8 9.75a.5.5 0 0 1 1 0c0 .924.808 1.25 1.5 1.25s1.5-.326 1.5-1.25c0-.608-.633-.89-1.676-1.281C9.281 8.078 8 7.6 8 6.25A2.281 2.281 0 0 1 10.5 4A2.281 2.281 0 0 1 13 6.25a.5.5 0 0 1-1 0C12 5.328 11.192 5 10.5 5S9 5.328 9 6.25c0 .609.633.89 1.676 1.282C11.719 7.923 13 8.4 13 9.75ZM6.75 4h-4.5a.5.5 0 1 0 0 1H4v6.5a.5.5 0 1 0 1 0V5h1.75a.5.5 0 1 0 0-1Z"/>`),
		g.Group(children),
	)
}