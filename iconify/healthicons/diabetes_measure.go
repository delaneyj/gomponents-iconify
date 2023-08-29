package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiabetesMeasure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M23.967 22.147c2.24-.002 3.999-1.68 3.997-3.813c-.002-1.98-4.005-6.854-4.005-6.854s-3.997 4.727-3.995 6.86c.002 2.134 1.763 3.809 4.003 3.807Zm.015 17.936a2 2 0 1 0-.003-4a2 2 0 0 0 .003 4Z"/><path fill-rule="evenodd" d="M12.952 4.092a3 3 0 0 0-2.997 3.002l.02 25a3 3 0 0 0 3.003 2.998l4.75-.004a1.26 1.26 0 0 0 .25-.026l.007 9.025l12-.01l-.007-9.024c.08.016.164.024.25.024l4.75-.003a3 3 0 0 0 2.997-3.003l-.02-25a3 3 0 0 0-3.003-2.998l-22 .019Zm20.004 3.983l-18 .015l.014 18l18-.015l-.014-18ZM15.973 29.59a1 1 0 0 1 1-1l4-.004a1 1 0 0 1 .001 2l-4 .003a1 1 0 0 1-1-.999Zm11-1.009a1 1 0 1 0 .001 2l4-.003a1 1 0 1 0-.002-2l-4 .003Zm1.003 4.5l-8 .006l.008 9l8-.007l-.008-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}