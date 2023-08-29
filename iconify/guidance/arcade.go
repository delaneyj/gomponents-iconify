package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arcade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 15.5v-8m0 0a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm-9.5 11v.219a6 6 0 0 1-1.516 3.986L.5 23.25v.25h23v-.25l-.485-.545a6 6 0 0 1-1.515-3.986V18.5h-19Zm3 0v-3h13v3h-13Z"/>`),
		g.Group(children),
	)
}