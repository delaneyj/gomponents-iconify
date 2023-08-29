package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 128L807 453l89 59v320L660 674l-30 45l266 177v128L571 807l-59 89H0l217-325l-89-59V384l148 99l29-45l-177-118V0l325 217l59-89h512z"/>`),
		g.Group(children),
	)
}