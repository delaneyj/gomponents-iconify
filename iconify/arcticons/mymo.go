package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mymo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsMymo0" d="M4.5 24.356v-3.199c0-2.014 1.45-3.634 3.25-3.634s3.248 1.62 3.249 3.634m0 3.199l.014-3.199c.014-2.014 1.464-3.634 3.264-3.634s3.249 1.62 3.249 3.634l-.006 2.625c-.003 1.44 1.443 3.061 3.243 3.061s3.229-1.621 3.216-4.132m6.507-1.554c0-2.014-1.449-3.634-3.249-3.634s-3.25 1.62-3.25 3.634l-.007 2.02l-.021 3.666c0 2.013-1.45 3.634-3.25 3.634c-.646 0-1.248-.21-1.753-.57m18.052-5.837c-.003-1.44 1.443-3.06 3.243-3.06s3.249 1.62 3.25 3.49c-.001 1.869-1.45 3.49-3.25 3.49s-3.246-1.62-3.244-3.06l.007-3.773c.003-2.014-1.45-3.634-3.25-3.634s-3.249 1.62-3.263 3.634l-.015 3.2"/></defs><use href="#arcticonsMymo0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsMymo0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}