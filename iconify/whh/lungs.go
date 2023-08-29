package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lungs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024q-13 0-38-16.5T861.5 966T800 928q-22-11-50-23t-45-19.5t-33.5-16.5t-24-18t-7.5-19V542L512 414L384 542v290q0 10-7.5 19t-24 18t-33.5 16.5t-45 19.5t-50 23q-26 13-61.5 38t-60.5 41.5t-38 16.5q-43 0-53.5-60T0 768q0-180 33-299t95-213q43-64 100-96t92-32q27 0 45.5 19t18.5 45v162l64-64V64q0-26 19-45t45.5-19t45 19T576 64v226l64 64V192q0-26 19-45t45-19q35 0 92 32t100 96q62 94 95 213t33 299q0 47-.5 70t-2 59.5t-5.5 54t-11 37t-18 27.5t-27 8z"/>`),
		g.Group(children),
	)
}