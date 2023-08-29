package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opensource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600.073 17.426C268.728 17.426 0 286.008 0 617.353c0 260.49 166.117 482.172 398.146 565.076L546.04 791.313c-74.269-23.013-128.273-92.136-128.273-173.96c0-100.646 81.661-182.307 182.308-182.307c100.646 0 182.16 81.66 182.16 182.307c0 81.888-53.952 151.147-128.273 174.106l147.896 391.115C1033.938 1099.72 1200 877.909 1200 617.353c0-331.345-268.581-599.927-599.927-599.927z"/>`),
		g.Group(children),
	)
}