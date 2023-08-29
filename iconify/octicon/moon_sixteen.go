package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.598 1.591a.749.749 0 0 1 .785-.175a7.001 7.001 0 1 1-8.967 8.967a.75.75 0 0 1 .961-.96a5.5 5.5 0 0 0 7.046-7.046a.75.75 0 0 1 .175-.786Zm1.616 1.945a7 7 0 0 1-7.678 7.678a5.499 5.499 0 1 0 7.678-7.678Z"/>`),
		g.Group(children),
	)
}