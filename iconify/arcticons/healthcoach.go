package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Healthcoach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.6V10.4a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27.2a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.004 34.042h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v1.3a2.006 2.006 0 0 1-2 2ZM8.033 13.958v8m0-3.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v3.3m23.934-8v8m0-3.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v3.3m-5.497 4.084v8m0-3.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v3.3m-6.152-19.084v7m-1.1-5.3h2.1m-14.934 4.3a1.936 1.936 0 0 1-1.7 1h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2v.7h-4m13.644-5.4v7a.945.945 0 0 0 1 1h.3M12.366 33.042a1.935 1.935 0 0 1-1.7 1h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a1.935 1.935 0 0 1 1.7 1m19.466 3.3a1.936 1.936 0 0 1-1.7 1h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a1.936 1.936 0 0 1 1.7 1m-6.497-9.784a2.006 2.006 0 0 1-2 2h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2m0 3.3v-5.3m.158 15.384a2.006 2.006 0 0 1-2 2h0a2.006 2.006 0 0 1-2-2v-1.3a2.006 2.006 0 0 1 2-2h0a2.006 2.006 0 0 1 2 2m0 3.3v-5.3"/>`),
		g.Group(children),
	)
}