package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.881 8.55v30.76c0 1.978 1.101 3.19 3.303 3.19h20.191c7.474 0 13.589-15.873 3.28-20.357c7.592-4.376 4.99-16.631-2.431-16.631c-5.067 0-21.529-.012-21.529-.012c-2.045 0-2.814.968-2.814 3.05h0Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="17.132" cy="19.712" rx="3.644" ry="3.639"/><ellipse cx="27.287" cy="19.712" rx="3.644" ry="3.639"/><ellipse cx="27.287" cy="29.944" rx="3.644" ry="3.639"/><ellipse cx="17.132" cy="29.944" rx="3.644" ry="3.639"/></g>`),
		g.Group(children),
	)
}