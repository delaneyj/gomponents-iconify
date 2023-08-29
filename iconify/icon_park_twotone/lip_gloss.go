package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LipGloss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLipGloss0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M24 12H8s0 4 1 14s3 18 3 18h8s2-8 3-18s1-14 1-14Zm18 22H28s0 2 1 5s2.5 5 2.5 5h7s1.5-2 2.5-5s1-5 1-5Z"/><path d="M35 34V13"/><path fill="#555" stroke-linejoin="round" d="M31 7h8l-2 6h-4l-2-6ZM11 6h10v6H11z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLipGloss0)"/>`),
		g.Group(children),
	)
}