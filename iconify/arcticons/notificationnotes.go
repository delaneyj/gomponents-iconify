package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notificationnotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.48 5.46a3.85 3.85 0 0 1 3 1.23l2.87 2.91a4 4 0 0 1 0 5.59L36 20.54L27.5 12l5.35-5.35a4.13 4.13 0 0 1 2.63-1.23ZM27.5 12l8.5 8.5l-14.69 14.73L17 39.56a1.88 1.88 0 0 1-.76.53l-9 2.41a1.37 1.37 0 0 1-1.68-1.68l2.4-9a1.89 1.89 0 0 1 .55-.77l1.65-1.65L27.5 12Z"/>`),
		g.Group(children),
	)
}