package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnAround(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm3 6h-6c-.929 0-1.393 0-1.784.038a8 8 0 0 0-7.178 7.178C12 27.607 12 28.07 12 29h24c0-.929 0-1.393-.038-1.784a8 8 0 0 0-7.178-7.178C28.393 20 27.93 20 27 20Z"/><path d="M41 26.784c1.902 1.224 3 2.669 3 4.216c0 4.418-8.954 8-20 8S4 35.418 4 31c0-1.547 1.098-2.992 3-4.216"/><path d="m19 34l5 5l-5 5"/></g>`),
		g.Group(children),
	)
}