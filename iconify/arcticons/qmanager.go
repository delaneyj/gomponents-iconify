package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.62 11.15h30.76v20H8.62Zm21 24.23h9.15a4.78 4.78 0 0 0 4.78-4.79V11.65a4.78 4.78 0 0 0-4.78-4.79H9.28a4.78 4.78 0 0 0-4.78 4.79v18.94a4.78 4.78 0 0 0 4.78 4.79h9.15v2.73h-3.58v3H33v-3h-3.43Z"/>`),
		g.Group(children),
	)
}