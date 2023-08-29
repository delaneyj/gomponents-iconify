package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForbiddenTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 12C.5 5.649 5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5S.5 18.351.5 12Zm7.342-6.836a8 8 0 0 1 10.994 10.994a67.213 67.213 0 0 1-2.216-1.829c-.883-.77-2.912-2.742-3.56-3.39c-.647-.647-2.62-2.676-3.39-3.56a67.247 67.247 0 0 1-1.828-2.215ZM5.164 7.842a8 8 0 0 0 10.994 10.994a67.213 67.213 0 0 0-1.829-2.216c-.77-.883-2.742-2.912-3.39-3.56c-.647-.647-2.676-2.62-3.56-3.39a67.226 67.226 0 0 0-2.215-1.828Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}