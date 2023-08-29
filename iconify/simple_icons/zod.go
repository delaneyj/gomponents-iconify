package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.088 2.477L24 7.606L12.521 20.485l-.925 1.038L0 7.559l5.108-5.082h13.98Zm-17.434 5.2l6.934-4.003H5.601L1.619 7.636l.035.041Zm12.117-4.003L3.333 9.7l2.149 2.588l10.809-6.241l-.2-.346l2.851-1.646l-.365-.381h-4.806Zm7.52 2.834L8.257 14.034h5.101v-.4h3.667l5.346-5.998l-1.08-1.128Zm-7.129 10.338H9.268l2.36 2.843l2.534-2.843Z"/>`),
		g.Group(children),
	)
}