package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mozc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.732 12.133s-1.747 10.937.152 20.165"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.39 16.614s5.013 1.48 14.81-1.481m-4.29 17.886s9.797-.456 8.506-7.253s-10.367-4.557-10.367-4.557s-9.152 1.215-9.494 8.506c-.23 4.904 3.975 2.552 6.947-.093a24.077 24.077 0 0 0 3.013-3.253a19.537 19.537 0 0 0 3.434-7.622"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.772 39.423a16.702 16.702 0 1 1 5.993-4.227"/><circle cx="35.365" cy="39.247" r="4.253" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}