package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gittouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 8.62l-4.96-2.86l-14.51 8.38l-14.6-8.43L4.5 8.55l19.56 11.3L43.5 8.62z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 8.55v5.7l14.6 8.43v16.75l4.96 2.86V19.85L4.5 8.55z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 8.62v5.7l-14.48 8.36v16.75l-4.96 2.86V19.85L43.5 8.62z"/>`),
		g.Group(children),
	)
}