package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modxalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 128L807 453l89 59v320L660 674l-30 45l266 177v128L571 807l-59 89H0l217-325l-89-59V384l148 98l29-44l-177-118V0l325 217l59-89h512zM268 606L117 832h363l39-59zm59-89l251 167l30-44l-251-167zm429-99l151-226H544l-39 59z"/>`),
		g.Group(children),
	)
}