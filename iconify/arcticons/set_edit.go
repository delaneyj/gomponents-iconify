package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.215 3.5l-1.004 5.172a15.778 15.778 0 0 0-4.6 2.647L8.64 9.606l-4.788 8.288l3.98 3.46a14.976 14.976 0 0 0 0 5.297l-3.98 3.459l4.789 8.284l4.964-1.71a15.641 15.641 0 0 0 4.605 2.643l1.005 5.173h9.569l.985-5.12a17.994 17.994 0 0 0 4.62-2.7l4.969 1.714l4.784-8.284l-3.976-3.463a14.746 14.746 0 0 0 0-5.289l3.98-3.467l-4.789-8.284l-4.964 1.71a15.689 15.689 0 0 0-4.605-2.643l-1.005-5.172Zm4.698 14.206a6.376 6.376 0 0 1 0 12.75h0a6.375 6.375 0 0 1-6.376-6.374v-.001a6.375 6.375 0 0 1 6.376-6.375Z"/>`),
		g.Group(children),
	)
}