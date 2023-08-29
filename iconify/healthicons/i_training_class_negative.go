package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ITrainingClassNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsITrainingClassNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM7 6H6v27h23.387v-2H8V8h27v3h2V6H7Zm32 10a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm2.03 6.496c-.672-.902-1.7-1.485-2.999-1.485h-6a1.72 1.72 0 0 0-.926.283c-.155.1-.284.212-.376.298a6.286 6.286 0 0 0-.543.602a24.45 24.45 0 0 0-1.08 1.475c-.336.49-.655.978-.898 1.354h-4.177a1.5 1.5 0 0 0 0 3h5c.518 0 1-.267 1.274-.707v-.002l.006-.007l.02-.032l.077-.123a61.361 61.361 0 0 1 .592-.92V40.5a1.5 1.5 0 0 0 2.974.276L35.432 33h.297l1.291 7.747A1.5 1.5 0 0 0 40 40.5v-9.842c.782-.617 1.288-1.517 1.59-2.388a7.64 7.64 0 0 0 .396-2.92c-.06-.964-.326-2.01-.956-2.854Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsITrainingClassNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}