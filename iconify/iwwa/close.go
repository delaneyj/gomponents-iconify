package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Close(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20.725 20L34.493 6.216a.516.516 0 0 0 0-.726a.526.526 0 0 0-.727.001L20 19.274L6.232 5.49a.528.528 0 0 0-.726 0a.515.515 0 0 0 0 .725L19.275 20L5.507 33.784a.513.513 0 1 0 .725.725L20 20.726L33.767 34.51a.513.513 0 0 0 .726-.725L20.725 20z"/>`),
		g.Group(children),
	)
}