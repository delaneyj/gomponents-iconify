package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bydesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBydesign0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 7.282L19.556 4v12.718L8 20V7.282Zm0 18.393l11.556-3.282v12.718L8 38.393V25.675Zm16.889-4.599l11.555-2.854v22.924L24.89 44V21.076Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBydesign0)"/>`),
		g.Group(children),
	)
}