package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaseDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 91q18 0 30.5 12.5T427 133v235q0 18-12.5 30.5T384 411H43q-18 0-30.5-12.5T0 368V133q0-17 12.5-29.5T43 91h85V48l43-43h85l43 43v43h85zM171 48v43h85V48h-85zm42 320l107-107h-64v-85h-85v85h-64z"/>`),
		g.Group(children),
	)
}