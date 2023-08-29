package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M58.25 2H5.75C3.875 2 2 3.875 2 5.75v42.188c0 1.875 1.875 3.75 3.75 3.75H24.5v3.75c0 2.604-3.378 3.75-7.5 3.75V62h30v-2.813c-4.123 0-7.5-1.146-7.5-3.75v-3.75h18.75c1.875 0 3.75-1.875 3.75-3.75V5.75C62 3.875 60.125 2 58.25 2zM60 47.938c0 .769-.982 1.75-1.75 1.75H5.75c-.769 0-1.75-.981-1.75-1.75v-3.75h56v3.75z"/><circle cx="32" cy="47.938" r=".938" fill="currentColor"/>`),
		g.Group(children),
	)
}