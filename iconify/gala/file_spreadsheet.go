package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSpreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaFileSpreadsheet0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileSpreadsheet1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileSpreadsheet2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileSpreadsheet3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileSpreadsheet4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileSpreadsheet5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileSpreadsheet6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileSpreadsheet7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileSpreadsheet8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileSpreadsheet9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileSpreadsheeta" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileSpreadsheetb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileSpreadsheetc" stroke-linejoin="round" stroke-width="16" d="m 64,192 h 96"/><path id="galaFileSpreadsheetd" stroke-linejoin="round" stroke-width="16" d="m 64,160 h 96"/><path id="galaFileSpreadsheete" stroke-linejoin="round" stroke-width="16" d="m 80,144 v 64"/><path id="galaFileSpreadsheetf" stroke-linejoin="round" stroke-width="16" d="m 112,144 v 64"/><path id="galaFileSpreadsheetg" stroke-linejoin="round" stroke-width="16" d="m 144,144 v 64"/></g>`),
		g.Group(children),
	)
}