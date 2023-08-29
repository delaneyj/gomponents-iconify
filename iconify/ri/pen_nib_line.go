package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.596 1.037l6.347 6.346a.5.5 0 0 1-.277.848l-1.474.23l-5.656-5.657l.212-1.485a.5.5 0 0 1 .848-.282ZM4.595 20.148c3.722-3.331 7.995-4.328 12.643-5.52l.446-4.018l-4.297-4.297l-4.018.446c-1.192 4.648-2.189 8.92-5.52 12.643l-1.395-1.395c2.829-3.3 3.89-6.953 5.303-13.081l6.364-.707l5.657 5.657l-.707 6.363c-6.128 1.415-9.782 2.475-13.081 5.304l-1.395-1.395Zm5.284-6.03a2 2 0 1 1 2.828-2.828a2 2 0 0 1-2.828 2.828Z"/>`),
		g.Group(children),
	)
}