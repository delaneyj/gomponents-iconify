package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 277v-42h43v42h-43zm0-192h43v107h-43V85zM170.5 21Q241 21 291 71t50 121t-50 121t-120.5 50T50 313T0 192T50 71t120.5-50zm0 214q17.5 0 30-12.5T213 192t-12.5-30.5t-30-12.5t-30 12.5T128 192t12.5 30.5t30 12.5z"/>`),
		g.Group(children),
	)
}