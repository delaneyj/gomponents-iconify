package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 85.9 0 192c0 75 57.5 139.8 141.1 171.4L85.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S397.4 0 256 0zm117.3 256L320 309.3l-64-64l-64 64l-53.3-53.3l64-64l-64-64L192 74.7l64 64l64-64l53.3 53.3l-64 64l64 64z"/>`),
		g.Group(children),
	)
}