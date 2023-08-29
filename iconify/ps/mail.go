package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21zm-40 43L256 166L83 64h346zM43 320V90l202 121q2 2 11 2t11-2L469 90v230H43z"/>`),
		g.Group(children),
	)
}