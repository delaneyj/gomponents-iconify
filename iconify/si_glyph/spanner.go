package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.838 11.784l5.906-5.905c1.172.605 2.567.493 3.463-.402c.69-.691.924-1.682.716-2.638l-1.522 1.519l-1.356.266l-1.641-1.625l.282-1.396l1.509-1.49c-.955-.208-1.947.023-2.638.714c-.896.896-1.008 2.29-.402 3.464l-5.906 5.906c-1.173-.605-2.568-.492-3.465.402c-.688.691-.922 1.682-.715 2.637l1.523-1.519l1.355-.265l1.643 1.625l-.284 1.396l-1.509 1.49c.955.207 1.947-.023 2.637-.714c.896-.895 1.009-2.291.404-3.465z"/>`),
		g.Group(children),
	)
}