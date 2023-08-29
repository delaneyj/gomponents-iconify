package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lawnchairsettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.8 3a1.06 1.06 0 0 0-1 .88L18 9.45a15.4 15.4 0 0 0-3.55 2.07L9.19 9.41a1 1 0 0 0-1.28.45l-4.2 7.27A1.06 1.06 0 0 0 4 18.48L8.39 22a17.67 17.67 0 0 0-.15 2a17.35 17.35 0 0 0 .15 2.05L4 29.52a1.06 1.06 0 0 0-.25 1.35l4.2 7.27a1 1 0 0 0 1.28.46l5.23-2.12A15.71 15.71 0 0 0 18 38.55l.79 5.57a1.06 1.06 0 0 0 1 .88h8.4a1.05 1.05 0 0 0 1-.88l.81-5.57a15.4 15.4 0 0 0 3.55-2.07l5.25 2.12a1 1 0 0 0 1.28-.46l4.2-7.27a1.06 1.06 0 0 0-.28-1.35l-4.4-3.47a17.35 17.35 0 0 0 .14-2.05a17.5 17.5 0 0 0-.14-2l4.4-3.52a1.06 1.06 0 0 0 .25-1.35l-4.2-7.27a1 1 0 0 0-1.28-.45l-5.23 2.11A15.71 15.71 0 0 0 30 9.45l-.79-5.57a1.07 1.07 0 0 0-1-.88Zm-4.03 12.85l16.31 16.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.87 23.95l-8.4 8.07l17.06-5.89"/>`),
		g.Group(children),
	)
}