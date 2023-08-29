package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampaignmonitorIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#7856FF" d="M254.234 4.152a9.725 9.725 0 0 0-13.565-2.381L1.737 169.047a9.705 9.705 0 0 0 7.968 4.148l.01.01h236.55c5.379 0 9.735-4.366 9.735-9.745V9.57a9.725 9.725 0 0 0-1.766-5.437M15.312 1.76A9.735 9.735 0 0 0 0 9.57v154.177l113.75-93.18L15.321 1.742"/>`),
		g.Group(children),
	)
}