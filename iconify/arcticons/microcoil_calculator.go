package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrocoilCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="27.881" cy="35.257" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.759" ry="3.879" transform="rotate(-44.99 27.881 35.257)"/><ellipse cx="24.792" cy="32.936" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.759" ry="3.879" transform="rotate(-44.99 24.792 32.936)"/><ellipse cx="21.703" cy="30.827" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.759" ry="3.879" transform="rotate(-44.99 21.703 30.827)"/><ellipse cx="18.614" cy="28.486" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.759" ry="3.879" transform="rotate(-44.99 18.614 28.486)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.214 28.884c-3.03 3.03-6.66 4.494-8.176 2.98s-.287-5.2 2.743-8.23a13.809 13.809 0 0 1 3.453-2.556m.104-.052c1.498-.79 2.733-1.695 2.975-3.613L19.01 4.5m19.6 8.704c-.51 6.68-1.901 17.083-1.532 20.037c.234 1.812-1.031 4.534-3.365 6.867c-3.03 3.03-6.715 4.258-8.23 2.743s-.308-4.864 2.722-7.893"/>`),
		g.Group(children),
	)
}