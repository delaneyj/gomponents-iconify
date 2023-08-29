package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adsilence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16.57v14.596h9.709l12.152 9.379V6.63l-12.02 9.94H4.5Zm39-10.468L7.802 41.898m28.086-27.847a12.089 12.089 0 0 1 0 19.805"/>`),
		g.Group(children),
	)
}