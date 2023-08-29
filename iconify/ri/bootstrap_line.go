package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BootstrapLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 11v2s2 0 2 4v1a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-1c0-4 2-4 2-4v-2s-2 0-2-4V6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v1c0 4-2 4-2 4Zm19.674 1a2.37 2.37 0 0 0-.088.086C19.578 13.093 19 14.692 19 17v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-1c0-2.308-.578-3.907-1.586-4.914A3.887 3.887 0 0 0 3.326 12a2.37 2.37 0 0 0 .088-.086C4.422 10.907 5 9.308 5 7V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v1c0 2.308.578 3.907 1.586 4.914l.088.086ZM8.5 7v10h5a3 3 0 0 0 1.615-5.529A3 3 0 0 0 12.5 7h-4Zm2 4V9h2a1 1 0 1 1 0 2h-2Zm0 2h3a1 1 0 1 1 0 2h-3v-2Z"/>`),
		g.Group(children),
	)
}