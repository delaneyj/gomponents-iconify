package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.687 1.777a1.5 1.5 0 0 1 2.629 0l5.499 9.999a1.5 1.5 0 0 1-1.314 2.223H2.502a1.5 1.5 0 0 1-1.314-2.223l5.499-9.999Zm1.752.482a.5.5 0 0 0-.876 0l-5.499 9.999a.5.5 0 0 0 .438.74h10.999a.5.5 0 0 0 .438-.74l-5.5-9.999Z"/>`),
		g.Group(children),
	)
}