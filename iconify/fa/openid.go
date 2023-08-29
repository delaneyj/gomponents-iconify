package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1086 0v1536l-272 128q-228-20-414-102t-293-208.5T0 1081q0-140 100.5-263.5t275-205.5T767 504v172q-217 38-356.5 150T271 1081q0 152 154.5 267T814 1493V133zm669 582l37 390l-525-114l147-83q-119-70-280-99V504q277 33 481 157z"/>`),
		g.Group(children),
	)
}