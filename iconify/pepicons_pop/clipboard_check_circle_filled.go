package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopClipboardCheckCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8.675 5.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H9.675a1 1 0 0 1-1-1V5.5Zm2 1v1.875h4.643V6.5h-4.643Z" clip-rule="evenodd"/><path d="M8 8v11h10V8h-1V6h2a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h2v2H8Z"/><path fill-rule="evenodd" d="M15.567 11.677a1 1 0 0 1 .256 1.39l-1.767 2.565a1.5 1.5 0 0 1-2.154.334l-1.515-1.176a1 1 0 0 1 1.226-1.58l1.097.851l1.466-2.128a1 1 0 0 1 1.391-.256Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopClipboardCheckCircleFilled0)"/></g>`),
		g.Group(children),
	)
}