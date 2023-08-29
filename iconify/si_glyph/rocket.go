package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 13c-.064.011-.103.027-.161.039c.701-2.108 1.133-5.049 1.133-7.175c0-4.315-2.932-5.837-3.965-5.837S5.04 1.548 5.04 5.864c0 2.125.432 5.063 1.132 7.171c-.057-.011-.11-.024-.173-.034c-.552 1.521-1.988 2.039-3.004 2.988l3.527.008c.276-.337.441-.72.476-1.127c.016.023.03.052.046.074h3.923c.028-.04.054-.089.082-.132c.031.412.192.802.467 1.143l3.471-.007C13.991 14.99 12.538 14.529 12 13zm-3-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a2 2 0 1 1 .001-4.001A2 2 0 0 1 9 6z"/>`),
		g.Group(children),
	)
}