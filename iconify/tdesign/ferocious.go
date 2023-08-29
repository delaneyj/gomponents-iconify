package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ferocious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5.769-3.866l3.464 2l-1 1.732l-3.464-2l1-1.732Zm11.464 1.732l-3.464 2l-1-1.732l3.464-2l1 1.732ZM16 12.791V17.1l-8.095-.771l-.092-1.978L16 12.791Z"/>`),
		g.Group(children),
	)
}