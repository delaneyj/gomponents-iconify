package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 7.25A3.25 3.25 0 0 1 5.25 4h13.5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20H5.25A3.25 3.25 0 0 1 2 16.75v-9.5ZM5.25 5.5A1.75 1.75 0 0 0 3.5 7.25v9.5c0 .966.784 1.75 1.75 1.75h13.5a1.75 1.75 0 0 0 1.75-1.75v-9.5a1.75 1.75 0 0 0-1.75-1.75H5.25ZM17.5 7.75a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5Zm.75 7.25a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Zm-.75-3.25a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5ZM5.75 7a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5A.75.75 0 0 0 5.75 7ZM5 15.75a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5ZM5.75 11a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}