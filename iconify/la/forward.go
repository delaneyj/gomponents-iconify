package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 5v7.375L6.594 6.187L5 5.032V26.97l1.594-1.157L15 19.625V27l1.594-1.188l12-9L29.656 16l-1.062-.813l-12-9zM7 8.969l8.406 6.187l1.125.844l-1.125.844L7 23.03zM17 9l9.313 7L17 23z"/>`),
		g.Group(children),
	)
}