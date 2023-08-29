package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sensu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 12L12 0L0 12l12 12l12-12zM12 3.197l4.418 4.418c-1.445-.386-2.93-.586-4.418-.586s-2.974.199-4.418.588L12 3.196zM8.069 16.87c1.19-.658 2.534-1.008 3.931-1.008s2.741.35 3.931 1.008L12 20.804L8.069 16.87zm9.509-1.647c-1.697-1.08-3.636-1.622-5.578-1.622s-3.881.542-5.578 1.622l-3.103-3.101C5.822 10.284 8.834 9.29 12 9.29s6.178.994 8.681 2.832l-3.103 3.101z"/>`),
		g.Group(children),
	)
}