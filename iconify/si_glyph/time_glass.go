package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12 2.026C12 .908 11.135 0 10.065 0h-3.12C5.877 0 5.01.908 5.01 2.026v3.536l2.021 2.427l-2.021 2.448v3.516c0 1.118.866 2.025 1.935 2.025h3.12c1.069 0 1.935-.907 1.935-2.025v-3.578L9.906 8.001L12 5.625V2.026zM11 5L9 7.989L11 11v3a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3l2-3.011L6 5.02V2a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v3z"/><path d="M10 14H7v-2l1.5-1.979L10 12v2zm0-10L8.521 5.979L7 4h3z"/></g>`),
		g.Group(children),
	)
}