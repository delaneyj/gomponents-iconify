package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccidentAndEmergencyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsAccidentAndEmergencyNegative0)"><path fill-rule="evenodd" d="M33 36H15v-2h18v2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20.489 16.33c-1.643.284-2.84 1.676-3.036 3.332L16 32h16l-1.456-12.355c-.194-1.648-1.38-3.037-3.014-3.323c-2.45-.429-4.492-.433-7.041.008Zm1.241 4.318a1 1 0 0 0-1.984-.248l-.905 7.238a1 1 0 1 0 1.984.248l.905-7.238Z" clip-rule="evenodd"/><path d="M13 24c0 .337.015.67.045 1H10v-2h3.045c-.03.33-.045.663-.045 1Zm2.012-6.344a10.98 10.98 0 0 0-1.001 1.732l-2.635-1.522l1-1.732l2.636 1.522Zm4.375-3.646c-.609.283-1.188.619-1.73 1.002l-1.523-2.636l1.732-1l1.521 2.635ZM24 13c-.337 0-.67.015-1 .045V10h2v3.045A11.17 11.17 0 0 0 24 13Zm6.344 2.012a10.989 10.989 0 0 0-1.732-1.001l1.522-2.635l1.732 1l-1.522 2.636Zm3.646 4.375a10.98 10.98 0 0 0-1.002-1.73l2.636-1.523l1 1.732l-2.635 1.521ZM35 24c0-.337-.015-.67-.045-1H38v2h-3.045c.03-.33.045-.663.045-1Z"/></g><defs><clipPath id="healthiconsAccidentAndEmergencyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}