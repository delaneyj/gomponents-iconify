package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePowerpointO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22zm384 1528V640H992q-40 0-68-28t-28-68V128H128v1536h1280zm-992-234v106h327v-106h-93v-167h137q76 0 118-15q67-23 106.5-87t39.5-146q0-81-37-141t-100-87q-48-19-130-19H416v107h92v555h-92zm353-280H650V882h120q52 0 83 18q56 33 56 115q0 89-62 120q-31 15-78 15z"/>`),
		g.Group(children),
	)
}