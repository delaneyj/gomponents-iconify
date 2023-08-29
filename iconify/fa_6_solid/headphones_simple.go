package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 80C141.1 80 48 173.1 48 288v104c0 13.3-10.7 24-24 24S0 405.3 0 392V288C0 146.6 114.6 32 256 32s256 114.6 256 256v104c0 13.3-10.7 24-24 24s-24-10.7-24-24V288c0-114.9-93.1-208-208-208zM80 352c0-35.3 28.7-64 64-64h16c17.7 0 32 14.3 32 32v128c0 17.7-14.3 32-32 32h-16c-35.3 0-64-28.7-64-64v-64zm288-64c35.3 0 64 28.7 64 64v64c0 35.3-28.7 64-64 64h-16c-17.7 0-32-14.3-32-32V320c0-17.7 14.3-32 32-32h16z"/>`),
		g.Group(children),
	)
}