package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JusticeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsJusticeNegative0)"><path d="m37.696 17l-3.236-5.177L30.901 17h6.795Zm-19.893 6.218l-3.839-5.933l-4.188 5.933h8.027Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM32.14 11.669l-.464.124l2.565-2.759a.995.995 0 0 1 1.12.455L40.053 17H42c0 4.418-3.358 8-7.5 8c-4.142 0-7.5-3.582-7.5-8h1.474l3.665-5.331ZM13.62 31.218c4.142 0 7.5-3.582 7.5-8h-.935l-4.6-7.108l5.842-1.567A3.01 3.01 0 0 0 23 15.829V38h-3v2h-6v2h20v-2h-6v-2h-3V15.83a3.001 3.001 0 0 0 2-2.782l4.672-1.254l2.564-2.758l-7.875 2.112A3.006 3.006 0 0 0 25 10.171V6h-2v4.17a3.004 3.004 0 0 0-1.97 2.409l-7.29 1.955a1 1 0 0 0-.557.39l-5.855 8.294H6.12c0 4.418 3.358 8 7.5 8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsJusticeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}