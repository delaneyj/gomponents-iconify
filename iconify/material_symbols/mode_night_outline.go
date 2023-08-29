package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeNightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 4h-.525q-.25 0-.475.05q1.425 1.65 2.212 3.688T11.5 12q0 2.225-.788 4.263T8.5 19.95q.225.05.475.05H9.5q3.325 0 5.663-2.337T17.5 12q0-3.325-2.337-5.663T9.5 4Zm0-2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T19.5 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T9.5 22q-1.325 0-2.588-.338T4.5 20.65Q6.825 19.3 8.163 17T9.5 12q0-2.7-1.338-5T4.5 3.35q1.15-.675 2.413-1.012T9.5 2Zm2 10Z"/>`),
		g.Group(children),
	)
}