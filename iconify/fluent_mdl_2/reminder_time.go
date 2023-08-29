package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReminderTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2029 1453l-557 558l-269-270l90-90l179 178l467-466l90 90zM1024 384v512h384v128H896V384h128zm-64 1408q16 0 32-1t32-2l114 114q-45 8-89 12t-89 5q-132 0-254-34t-230-97t-194-150t-150-195t-97-229T0 960q0-132 34-254t97-230t150-194t195-150t229-97T960 0q132 0 254 34t230 97t194 150t150 195t97 229t35 255q0 50-6 101l-168 169q46-133 46-270q0-115-30-221t-84-198t-130-169t-168-130t-199-84t-221-30q-115 0-221 30t-198 84t-169 130t-130 168t-84 199t-30 221q0 115 30 221t84 198t130 169t168 130t199 84t221 30z"/>`),
		g.Group(children),
	)
}