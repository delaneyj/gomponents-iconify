package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazeUtilities(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.27 28.377h0a2.167 2.167 0 0 1 0-3.754h0a.766.766 0 0 0 .28-1.046h0l-1.244-2.154a.766.766 0 0 0-1.046-.28h0a2.167 2.167 0 0 1-3.25-1.877h0a.766.766 0 0 0-.766-.766h-2.488a.766.766 0 0 0-.766.766h0a2.167 2.167 0 0 1-3.25 1.877h0a.766.766 0 0 0-1.046.28l-1.244 2.154a.766.766 0 0 0 .28 1.046h0a2.167 2.167 0 0 1 0 3.753h0a.766.766 0 0 0-.28 1.047h0l1.244 2.154a.766.766 0 0 0 1.046.28h0a2.167 2.167 0 0 1 3.25 1.877h0c0 .423.343.766.766.766h2.488a.766.766 0 0 0 .766-.766h0a2.167 2.167 0 0 1 3.25-1.876h0a.766.766 0 0 0 1.046-.28l1.244-2.155a.766.766 0 0 0-.28-1.046s0 0 0 0h0Z"/><path d="m24 23.5l.918 2.082L27 26.5l-2.082.918L24 29.5l-.918-2.082L21 26.5l2.082-.918L24 23.5z"/></g>`),
		g.Group(children),
	)
}