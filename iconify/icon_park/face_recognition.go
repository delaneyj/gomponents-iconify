package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M33 4.99976H41C42.1046 4.99976 43 5.89519 43 6.99976V14.9998M43 32.9998V40.9998C43 42.1043 42.1046 42.9998 41 42.9998H33M15 42.9998H7C5.89543 42.9998 5 42.1043 5 40.9998V32.9998M5 14.9998V6.99976C5 5.89519 5.89543 4.99976 7 4.99976H15"/><path d="M24 38C30.6274 38 36 31.732 36 24C36 16.268 30.6274 10 24 10C17.3726 10 12 16.268 12 24C12 31.732 17.3726 38 24 38Z"/><path stroke-linecap="round" d="M6 24H42"/><path stroke-linecap="round" d="M20.0693 30.1057C21.3372 31.0429 22.6473 31.5115 23.9996 31.5115C25.3519 31.5115 26.698 31.0429 28.0378 30.1057"/></g>`),
		g.Group(children),
	)
}