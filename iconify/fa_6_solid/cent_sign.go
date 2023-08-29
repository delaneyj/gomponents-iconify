package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 0c17.7 0 32 14.3 32 32v34.7c30.9 5.2 59.2 17.7 83.2 35.8c14.1 10.6 17 30.7 6.4 44.8s-30.7 17-44.8 6.4C279.4 137.5 252.9 128 224 128c-70.7 0-128 57.3-128 128s57.3 128 128 128c28.9 0 55.4-9.5 76.8-25.6c14.1-10.6 34.2-7.8 44.8 6.4s7.8 34.2-6.4 44.8c-24 18-52.4 30.6-83.2 35.8V480c0 17.7-14.3 32-32 32s-32-14.3-32-32v-34.7C101.2 430.1 32 351.1 32 256S101.2 81.9 192 66.7V32c0-17.7 14.3-32 32-32z"/>`),
		g.Group(children),
	)
}