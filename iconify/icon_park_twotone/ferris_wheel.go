package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FerrisWheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFerrisWheel0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="m9 44l15-22m15 22L24 22"/><path stroke="#fff" stroke-width="4" d="M9.132 24A15.14 15.14 0 0 1 9 22a14.95 14.95 0 0 1 1.5-6.546m.656 14.299A15.018 15.018 0 0 0 21 36.7m6 0a15.02 15.02 0 0 0 9.911-7.06M38.868 24c.087-.654.132-1.322.132-2a14.94 14.94 0 0 0-1.5-6.546M27 7.3a14.955 14.955 0 0 1 7 3.52M21 7.3a14.955 14.955 0 0 0-7.614 4.101"/><circle cx="10" cy="27" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="24" cy="37" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="24" cy="7" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="12" cy="13" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="36" cy="13" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="38" cy="27" r="3" fill="#555" stroke="#fff" stroke-width="4"/><circle cx="24" cy="22" r="4" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 44h8m20 0h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFerrisWheel0)"/>`),
		g.Group(children),
	)
}