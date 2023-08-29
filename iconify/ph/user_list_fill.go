package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserListFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 80a8 8 0 0 1 8-8h96a8 8 0 0 1 0 16h-96a8 8 0 0 1-8-8Zm104 40h-96a8 8 0 0 0 0 16h96a8 8 0 0 0 0-16Zm0 48h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm-138.71-26a48 48 0 1 0-58.58 0c-20.62 8.73-36.87 26.3-42.46 48A8 8 0 0 0 16 200h128a8 8 0 0 0 7.75-10c-5.59-21.71-21.84-39.28-42.46-48Z"/>`),
		g.Group(children),
	)
}