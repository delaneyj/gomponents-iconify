package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 14.95l2.775 2.1q.3.2.6.013t.175-.538L14.5 13.05l2.725-1.95q.3-.225.175-.563t-.475-.337H13.6l-1.125-3.65Q12.35 6.2 12 6.2t-.475.35L10.4 10.2H7.075q-.35 0-.475.338t.175.562L9.5 13.05l-1.05 3.475q-.125.35.175.538t.6-.013L12 14.95ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-10Zm0 8q3.325 0 5.663-2.337T20 12q0-3.325-2.337-5.663T12 4Q8.675 4 6.337 6.337T4 12q0 3.325 2.337 5.663T12 20Z"/>`),
		g.Group(children),
	)
}