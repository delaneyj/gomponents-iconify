package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M8 22H7v-9h2v9H8Zm26.556 12.58c.332-.453.666-.845.944-1.152c.278.307.612.7.944 1.152c.842 1.15 1.556 2.522 1.556 3.784C38 39.856 36.845 41 35.5 41S33 39.856 33 38.364c0-1.262.714-2.635 1.556-3.784ZM40 25h1v4H29v-4h2v-2a3 3 0 0 0-3-3v-8h2a9 9 0 0 1 9 9v4h1ZM19 12h5v11H13V12h6Z"/>`),
		g.Group(children),
	)
}