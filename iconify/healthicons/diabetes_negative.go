package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiabetesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDiabetesNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm24.126 17.842l17.719-.015a2.128 2.128 0 1 1 .003 4.257h-.005c.086-.168.133-.35.133-.54c0-.828-.897-1.5-2.001-1.498c-1.105 0-2 .673-1.999 1.501c0 .19.047.372.134.54l-13.022.01v1.064l6.89-.005a2 2 0 0 1 2.001 1.998v.256a2 2 0 0 1-1.998 2.002l-6.889.006v1.064l5.779-.005a2 2 0 0 1 2.001 1.998v.256a2 2 0 0 1-1.998 2.002l-8 .007l.001 1.063l4.111-.003a2 2 0 0 1 2.002 1.998v.257a2 2 0 0 1-1.998 2.001l-7.445.006l-4.1.004c-6.326.005-11.46-5.12-11.465-11.447a11.456 11.456 0 0 1 7.117-10.612l12.883-5.273a2.308 2.308 0 0 1 1.834 4.235l-2.105.964c-.977.447-.658 1.91.417 1.909Zm15.859 14.204c1.68-.002 2.999-1.26 2.997-2.86c0-1.486-3.004-5.14-3.004-5.14s-2.997 3.545-2.996 5.145c.002 1.6 1.323 2.856 3.003 2.855Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDiabetesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}