package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m239.15 212.42l-56-112a8 8 0 0 0-14.31 0l-21.71 43.43A88 88 0 0 1 100 126.93A103.65 103.65 0 0 0 127.69 64H152a8 8 0 0 0 0-16H96V32a8 8 0 0 0-16 0v16H24a8 8 0 0 0 0 16h87.63A87.76 87.76 0 0 1 88 116.35a87.74 87.74 0 0 1-19-31a8 8 0 1 0-15.08 5.34A103.63 103.63 0 0 0 76 127a87.55 87.55 0 0 1-52 17a8 8 0 0 0 0 16a103.46 103.46 0 0 0 64-22.08a104.18 104.18 0 0 0 51.44 21.31l-26.6 53.19a8 8 0 0 0 14.31 7.16L140.94 192h70.11l13.79 27.58A8 8 0 0 0 232 224a8 8 0 0 0 7.15-11.58ZM148.94 176L176 121.89L203.05 176Z"/>`),
		g.Group(children),
	)
}