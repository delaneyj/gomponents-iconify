package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CediSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 32c0-17.7-14.3-32-32-32s-32 14.3-32 32v34.7C101.2 81.9 32 160.9 32 256s69.2 174.1 160 189.3V480c0 17.7 14.3 32 32 32s32-14.3 32-32v-34.7c30.9-5.2 59.2-17.7 83.2-35.8c14.1-10.6 17-30.7 6.4-44.8s-30.7-17-44.8-6.4c-13.2 9.9-28.3 17.3-44.8 21.6V132c16.4 4.2 31.6 11.6 44.8 21.6c14.1 10.6 34.2 7.8 44.8-6.4s7.8-34.2-6.4-44.8c-24-18-52.4-30.6-83.2-35.8V32zm-64 100v248c-55.2-14.2-96-64.3-96-124s40.8-109.8 96-124z"/>`),
		g.Group(children),
	)
}