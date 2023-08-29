package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FootprintFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h5.5v1.25a2.75 2.75 0 1 1-5.5 0V18ZM8 6.12c2 0 3 2.88 3 4.88c0 1-.5 2-1 3.5L9.5 16H4c0-1-.5-2.5-.5-5S5.498 6.12 8 6.12Zm12.054 7.978l-.217 1.231a2.75 2.75 0 1 1-5.416-.955l.216-1.23l5.417.954ZM18.178 1.705c2.464.434 4.018 3.124 3.584 5.586c-.434 2.462-1.187 3.853-1.36 4.838l-5.417-.955l-.232-1.565c-.232-1.564-.55-2.635-.377-3.62c.347-1.97 1.832-4.632 3.802-4.284Z"/>`),
		g.Group(children),
	)
}