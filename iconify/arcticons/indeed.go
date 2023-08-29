package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.758 27.267V40.53c0 3.958 6.964 3.96 6.964 0V27.267a7.719 7.719 0 0 1-6.964 0Z"/><ellipse cx="26.468" cy="16.146" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.17" ry="4.139"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.518 22.771c2.33-9.728 11.661-25.061 26.964-14.96"/>`),
		g.Group(children),
	)
}