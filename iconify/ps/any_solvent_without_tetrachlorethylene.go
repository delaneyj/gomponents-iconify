package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnySolventWithoutTetrachlorethylene(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M337 220q0 35-21 53q-22 19-62 19h-19v77h-45V156h68q37 0 60 17q19 12 19 47zm-102 32h15q20 0 32-9q10-7 10-23q0-17-8-24q-14-8-28-8h-21v64zM256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469z"/>`),
		g.Group(children),
	)
}