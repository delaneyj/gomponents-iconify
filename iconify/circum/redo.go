package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.061 13.67A7.958 7.958 0 0 0 16.2 19.74a8.061 8.061 0 0 0 3.77-6.77a.5.5 0 0 0-1 0a6.976 6.976 0 0 1-11 5.7a6.969 6.969 0 0 1 1-11.97a10.075 10.075 0 0 1 4.64-.69v1.45a.5.5 0 0 0 .81.39l2.47-1.95a.5.5 0 0 0 0-.79L14.4 3.17a.5.5 0 0 0-.8.4v1.44c-.71-.01-1.43-.03-2.13.02a7.985 7.985 0 0 0-7.41 8.64Z"/>`),
		g.Group(children),
	)
}