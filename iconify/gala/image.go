package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaImage0"><g id="galaImage1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.233" transform="scale(3.77953)"><rect id="galaImage2" width="59.267" height="59.267" x="4.233" y="4.233" stroke-opacity="1" ry="8.467"/><path id="galaImage3" stroke-opacity="1" d="m 25.4,33.866667 c -12.7,0 -16.9333334,8.466665 -21.1666669,8.466665"/><path id="galaImage4" stroke-opacity="1" d="m 25.4,33.866667 c 12.7,0 17.582797,8.466665 25.399999,8.466665"/><path id="galaImage5" stroke-opacity="1" d="M 63.5,38.099999 C 48.683332,38.099998 46.566666,50.8 38.1,50.8"/><circle id="galaImage6" cx="46.567" cy="21.167" r="8.467"/></g></g>`),
		g.Group(children),
	)
}