package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.99L8.434 5.07L2.73 13.057L16 28.537l13.268-15.48l-5.704-7.987L16 3.99zm-2.53 2.381l-3.267 3.734l.227-3.298l3.04-.436zm5.06 0l3.04.434l.227 3.3l-3.268-3.734zM16 6.52L20.797 12h-9.594L16 6.52zM8.295 8.707L8.066 12H5.943l2.352-3.293zm15.41.002L26.057 12h-2.123l-.229-3.291zM6.175 14h2.208l3.09 6.182L6.176 14zm4.444 0h10.762L16 24.764L10.62 14zm12.998 0h2.207l-5.297 6.182L23.617 14z"/>`),
		g.Group(children),
	)
}