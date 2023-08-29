package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsRemote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 23q-.425 0-.713-.288T8 22V10q0-.425.288-.713T9 9h6q.425 0 .713.288T16 10v12q0 .425-.288.713T15 23H9Zm3-8.75q.525 0 .888-.363T13.25 13q0-.525-.363-.888T12 11.75q-.525 0-.888.363T10.75 13q0 .525.363.888t.887.362Zm-3.55-6.8l-1.4-1.4q1-1 2.275-1.525T12 4q1.4 0 2.675.525T16.95 6.05l-1.4 1.4q-.725-.725-1.637-1.088T12 6q-1 0-1.913.363T8.45 7.45Zm-2.8-2.8L4.2 3.2Q5.775 1.675 7.788.837T12 0q2.2 0 4.213.838T19.75 3.25l-1.4 1.4q-1.25-1.3-2.9-1.975T12 2q-1.8 0-3.45.675T5.65 4.65Z"/>`),
		g.Group(children),
	)
}