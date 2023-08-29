package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthyRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHealthyRecognition0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 33v6a3 3 0 0 0 3 3h6m18 0h6a3 3 0 0 0 3-3v-6m0-18V9a3 3 0 0 0-3-3h-6M6 15V9a3 3 0 0 1 3-3h6"/><path fill="#555" d="M34 24L24 34L14 24c-1-1-1.5-3 0-5s4.5-2 6-1s2 2 4 2s2.5-1 4-2s4.5-1 6 1s1 4 0 5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHealthyRecognition0)"/>`),
		g.Group(children),
	)
}