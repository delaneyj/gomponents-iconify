package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BnhBemobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.779 20.997l12.941-8.171l6.233 6.09m.269 3.284l9.723-6.4l7.259 6.145M5.506 38.05H42.5M5.945 30.498H42.5M31.307 15.976s-2.877-2.761-8.008-1.651M42.5 24.515a2.52 2.52 0 0 1-3.791 1.148a2.568 2.568 0 0 1-2.908-.007a2.568 2.568 0 0 1-2.907-.006a2.568 2.568 0 0 1-2.907-.007a2.568 2.568 0 0 1-2.907-.007a2.568 2.568 0 0 1-2.908-.007a2.568 2.568 0 0 1-2.907-.007a2.568 2.568 0 0 1-2.907-.007a2.568 2.568 0 0 1-2.908-.007a2.568 2.568 0 0 1-2.907-.007a2.568 2.568 0 0 1-2.908-.007v.001a2.917 2.917 0 0 1-4.135-1.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.299 5.5H40.7a1.799 1.799 0 0 1 1.8 1.799V40.7a1.799 1.799 0 0 1-1.799 1.799H7.3a1.799 1.799 0 0 1-1.8-1.798V7.3a1.799 1.799 0 0 1 1.799-1.8Z"/>`),
		g.Group(children),
	)
}