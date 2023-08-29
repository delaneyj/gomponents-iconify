package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zdfheute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.067 26.616v3.3a2 2 0 0 0 4 0v-3.3m0 3.3v2m-17.524-8v8m0-3.3a2 2 0 0 1 4 0v3.3m-3.996-16.05h4l-4 5.3h4m6.459 9.75a1.936 1.936 0 0 1-1.7 1a2.006 2.006 0 0 1-2-2v-1.3a2 2 0 0 1 4 0v.7h-4m14.31-4.4v6a.945.945 0 0 0 1 1h.3m-2.3-5.3h2.1m-11.41-8.75a2 2 0 0 0-4 0v1.3a2 2 0 0 0 4 0m0 2v-8m3.862 8v-6.6a1.367 1.367 0 0 1 1.4-1.4a1.557 1.557 0 0 1 1.4.6m-3.9 2.1h2.8m12.289 15.05a1.936 1.936 0 0 1-1.7 1a2.006 2.006 0 0 1-2-2v-1.3a2 2 0 1 1 4 0v.7h-4"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M9.543 34.834h4"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}