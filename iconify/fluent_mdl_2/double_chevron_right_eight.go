package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronRightEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m187 262l762 762l-762 762L6 1605l581-581L6 443l181-181zm1786 762l-762 762l-181-181l581-581l-581-581l181-181l762 762z"/>`),
		g.Group(children),
	)
}