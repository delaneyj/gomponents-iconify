package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CriticalCareNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsCriticalCareNegative0)" clip-rule="evenodd"><path d="M13 13a1 1 0 0 0-1 1v15.96a1 1 0 0 0 1 1h10v2.142h-4.538v2h11.076v-2H25v-2.143h10a1 1 0 0 0 1-1V14a1 1 0 0 0-1-1H13Zm9.907 2.694l-3.403 5.985h-5.658v1.987h6.833l1.619-2.847l1.891 6.548l4.43-5.633h4.53a1 1 0 0 0 1.005-.993a1 1 0 0 0-1.005-.994h-5.513l-2.602 3.31l-2.127-7.363Z"/><path d="M48 0H0v48h48V0ZM9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Z"/></g><defs><clipPath id="healthiconsCriticalCareNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}