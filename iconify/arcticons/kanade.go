package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.916 30.304v10.208M11.6 42.5h24.824m1.66-12.196v10.208M24.012 24.863v13.12m0-.003l6.821-8.024m-13.695 0l6.876 8.027M9.916 15.599v6.07l28.168-.028v-5.898m-2.072-.144V5.5h-2.414m-3.052 0H17.8m-2.964 0h-2.524v9.811M19.18 5.5l-6.868 6.819M29.107 5.5l6.905 6.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.331 13.485h2.553v2.553h-2.553zm10.808 0h2.553v2.553h-2.553z"/>`),
		g.Group(children),
	)
}