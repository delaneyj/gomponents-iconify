package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMicOffOutline0"><g id="evaMicOffOutline1"><g id="evaMicOffOutline2" fill="currentColor"><path d="M10 6a2 2 0 0 1 4 0v5a1 1 0 0 1 0 .16l1.6 1.59A4 4 0 0 0 16 11V6a4 4 0 0 0-7.92-.75L10 7.17Zm9 5a1 1 0 0 0-2 0a4.86 4.86 0 0 1-.69 2.48L17.78 15A7 7 0 0 0 19 11Zm-7 4h.16L8 10.83V11a4 4 0 0 0 4 4Zm8.71 4.29l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/><path d="M15 20h-2v-2.08a7 7 0 0 0 1.65-.44l-1.6-1.6A4.57 4.57 0 0 1 12 16a5 5 0 0 1-5-5a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`),
		g.Group(children),
	)
}