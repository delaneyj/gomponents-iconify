package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiabetesMeasureNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDiabetesMeasureNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm12.953 4.336a3 3 0 0 0-2.998 3.002l.02 25a3 3 0 0 0 3.003 2.998l4.75-.004c.086 0 .17-.01.25-.026l.006 7.025l.001 2l2-.001l8-.007l2-.001l-.001-2l-.006-7.025c.08.016.164.025.25.024l4.75-.003a3 3 0 0 0 2.997-3.003l-.02-25a3 3 0 0 0-3.002-2.997l-22 .018Zm20.003 3.983l-18 .015l.015 18l18-.015l-.015-18ZM15.974 29.833a1 1 0 0 1 .999-1l4-.004a1 1 0 1 1 .001 2l-4 .003a1 1 0 0 1-1-.999Zm10.999-1.009a1 1 0 1 0 .001 2l4-.003a1 1 0 0 0-.001-2l-4 .003Zm-3.006-6.433c2.24-.002 4-1.68 3.997-3.813c-.001-1.98-4.005-6.853-4.005-6.853s-3.996 4.726-3.995 6.86c.002 2.133 1.763 3.808 4.003 3.806Zm4.01 10.932l-8 .007l.007 9l8-.007l-.008-9Zm-3.995 7.004a2 2 0 1 0-.003-4a2 2 0 0 0 .003 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDiabetesMeasureNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}