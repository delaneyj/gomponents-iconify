package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCamerasTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19.0059 26.2759V37.0001H5"/><path d="M42.6205 21.2152L38.7568 20.1799L34.7544 27.3897L40.55 28.9426L42.6205 21.2152Z"/><path fill="#2F88FF" d="M38.7566 20.18L34.7542 27.3898L33.0118 30.0287L5 22.523L8.62347 9L40.499 17.541L38.7566 20.18Z"/></g>`),
		g.Group(children),
	)
}