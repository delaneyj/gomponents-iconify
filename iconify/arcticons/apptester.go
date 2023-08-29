package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apptester(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.495 13.352H41.85a1.65 1.65 0 0 1 1.624 1.943l-3.157 17.502a2.193 2.193 0 0 1-2.154 1.804l-23.585.047a2.283 2.283 0 1 1 0-4.567h2.693a2.283 2.283 0 0 0 0-4.566h-4.327a2.219 2.219 0 1 1 0-4.436h6.094"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.495 13.352a2.251 2.251 0 0 0-2.283 2.218a2.41 2.41 0 0 0 2.555 2.276c1.26.059 2.063.77 2.063 1.646a1.699 1.699 0 0 1-1.792 1.588"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m24.251 29.766l2.071-11.481h11.116l-2.072 11.48Z"/><circle cx="6.699" cy="32.448" r="2.199" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}