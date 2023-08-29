package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M330.5 85q-17.5 0-30-12.5t-12.5-30t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5zm-224 139q44.5 0 75.5 31t31 75.5t-31 75.5t-75.5 31T31 406T0 330.5T31 255t75.5-31zm0 181q30.5 0 52.5-22t22-52.5t-22-52.5t-52.5-22T54 278t-22 52.5T54 383t52.5 22zM230 192l47 49v132h-42V267l-69-60q-12-10-12-30q0-17 12-30l60-60q10-12 30-12q18 0 34 12l41 41q32 32 76 32v43q-64 0-108-45l-17-17zm175.5 32q44.5 0 75.5 31t31 75.5t-31 75.5t-75.5 31t-75.5-31t-31-75.5t31-75.5t75.5-31zm0 181q30.5 0 52.5-22t22-52.5t-22-52.5t-52.5-22t-52.5 22t-22 52.5t22 52.5t52.5 22z"/>`),
		g.Group(children),
	)
}