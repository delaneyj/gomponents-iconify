package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeQuarterCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 896q0 59-3 114t-11 109t-23 107t-38 108q-57 134-148 242t-206 184t-251 118t-280 42q-133 0-255-34t-230-96t-194-150t-150-195t-97-229T0 960q0-133 34-255t96-230t150-194t195-150t229-97T960 0h64v896h896zm-960 896q108 0 209-27t191-76t165-119t132-155t90-184t43-207H896V130q-108 8-207 42t-184 91t-155 131t-119 165t-76 191t-27 210q0 115 30 221t84 198t130 169t168 130t199 84t221 30z"/>`),
		g.Group(children),
	)
}