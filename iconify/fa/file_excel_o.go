package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcelO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22zm384 1528V640H992q-40 0-68-28t-28-68V128H128v1536h1280zm-979-234v106h281v-106h-75l103-161q5-7 10-16.5t7.5-13.5t3.5-4h2q1 4 5 10q2 4 4.5 7.5t6 8t6.5 8.5l107 161h-76v106h291v-106h-68l-192-273l195-282h67V768H828v107h74l-103 159q-4 7-10 16.5t-9 13.5l-2 3h-2q-1-4-5-10q-6-11-17-23L648 875h76V768H434v107h68l189 272l-194 283h-68z"/>`),
		g.Group(children),
	)
}