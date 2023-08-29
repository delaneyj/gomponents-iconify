package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16.016c1.245.529 2 1.223 2 1.984c0 1.657-3.582 3-8 3s-8-1.343-8-3c0-.76.755-1.456 2-1.984"/><path fill="currentColor" fill-rule="evenodd" d="M11.262 17.675L12 17l-.738.675zm1.476 0l.005-.005l.012-.014l.045-.05l.166-.186a38.19 38.19 0 0 0 2.348-2.957c.642-.9 1.3-1.92 1.801-2.933c.49-.99.885-2.079.885-3.086C18 4.871 15.382 2 12 2S6 4.87 6 8.444c0 1.007.395 2.096.885 3.086c.501 1.013 1.16 2.033 1.8 2.933a38.153 38.153 0 0 0 2.515 3.143l.045.05l.012.014l.005.005a1 1 0 0 0 1.476 0zM12 17l.738.674L12 17zm0-11a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}