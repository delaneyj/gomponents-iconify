package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tonality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm-1-2.05V4.05q-3.025.375-5.013 2.65T4 12q0 3.025 1.988 5.3T11 19.95Zm2 0q.75-.125 1.475-.338T15.85 19H13v.95ZM13 17h5.25q.2-.225.35-.475t.3-.525H13v1Zm0-3h6.75l.1-.5l.1-.5H13v1Zm0-3h6.95l-.1-.5l-.1-.5H13v1Zm0-3h5.9q-.15-.275-.3-.525T18.25 7H13v1Zm0-3h2.85q-.65-.4-1.375-.613T13 4.05V5Z"/>`),
		g.Group(children),
	)
}