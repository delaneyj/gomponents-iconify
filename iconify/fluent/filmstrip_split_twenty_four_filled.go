package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.75 2.75a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75ZM2 7.25A3.25 3.25 0 0 1 5.25 4h5v16h-5A3.25 3.25 0 0 1 2 16.75v-9.5ZM4.75 7a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5A.75.75 0 0 0 4.75 7ZM4 15.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-1.5 0ZM4.75 11a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Zm9 9V4h5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20h-5ZM18.5 7.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-1.5 0Zm.75 7.25a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75Zm-.75-3.25v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-1.5 0Z"/>`),
		g.Group(children),
	)
}