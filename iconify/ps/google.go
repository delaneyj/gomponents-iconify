package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M303 2H178Q110 2 70 35q-39 33-39 81q0 38 27.5 66.5T132 211q3 0 9-.5t10-.5q-6 14-6 23q0 16 18 41q-79 6-115 28q-47 27-47 74q0 36 33.5 61t96.5 25q74 0 116.5-34t42.5-80q0-20-8.5-36.5t-16-24.5t-25.5-24l-22-17q-16-11-16-26q0-12 17-29q17-13 25.5-21.5t17-26T270 105q0-45-43-82h37zm-53 370q0 29-23.5 47T161 437q-48 0-77.5-20.5T54 363q0-42 52-62q29-10 64-10q9 0 14 1q38 26 52 42t14 38zm-49-195q-15 17-41 17q-35 0-55-35.5T85 86q0-28 13-45q16-19 42-19q34 0 55.5 36.5T217 133q0 28-16 44z"/>`),
		g.Group(children),
	)
}