package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedRealityBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path fill-rule="evenodd" d="M14 7.25a.75.75 0 0 0-.75.75v8a.75.75 0 0 0 1.5 0v-3.25h.335l2.28 3.648a.75.75 0 1 0 1.271-.796l-1.853-2.965A2.751 2.751 0 0 0 16 7.25h-2Zm.75 1.5v2.5H16a1.25 1.25 0 1 0 0-2.5h-1.25Zm-5.534-.974a.75.75 0 0 0-1.432 0l-2.5 8a.75.75 0 0 0 1.432.448l.71-2.274h2.148l.71 2.274a.75.75 0 0 0 1.432-.448l-2.5-8Zm-1.32 4.674h1.21L8.5 10.515l-.605 1.935Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}