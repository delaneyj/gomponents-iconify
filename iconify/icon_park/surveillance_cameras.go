package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCameras(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 6H23H5"/><path d="M23 23V6"/><path fill="#2F88FF" d="M8.42498 19.5798L40.3005 28.1209L38.5581 30.7598L34.5557 37.9696L32.8133 40.6085L4.80151 33.1028L8.42498 19.5798Z"/><path d="M38.5583 30.7598L42.422 31.7951L40.3515 39.5225L34.5559 37.9696"/></g>`),
		g.Group(children),
	)
}