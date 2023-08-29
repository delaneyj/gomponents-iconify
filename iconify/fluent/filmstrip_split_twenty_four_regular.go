package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.75 2.75a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75ZM2 7.25A3.25 3.25 0 0 1 5.25 4h5v1.5h-5A1.75 1.75 0 0 0 3.5 7.25v9.5c0 .966.784 1.75 1.75 1.75h5V20h-5A3.25 3.25 0 0 1 2 16.75v-9.5ZM13.75 5.5V4h5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20h-5v-1.5h5a1.75 1.75 0 0 0 1.75-1.75v-9.5a1.75 1.75 0 0 0-1.75-1.75h-5Zm3.75 2.25a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5Zm.75 7.25a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Zm-.75-3.25a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5ZM5.75 7a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5A.75.75 0 0 0 5.75 7ZM5 15.75a.75.75 0 0 1 1.5 0v.5a.75.75 0 0 1-1.5 0v-.5ZM5.75 11a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}