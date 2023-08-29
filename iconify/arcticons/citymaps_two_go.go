package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CitymapsTwoGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.468 30.406a18.52 18.52 0 0 0-7.88 11.727"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 42.5V30.406M11.73 36h24.54"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.532 30.406a18.52 18.52 0 0 1 7.88 11.727"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.59 39.978a20 20 0 0 1 36.82 0M8.797 13v5.3a2.67 2.67 0 0 0 2.6 2.7a2.67 2.67 0 0 0 2.6-2.7V13m2.056 0v7a.945.945 0 0 0 1 1h.3m2.056-3.2a2 2 0 0 1 4 0V21m-4-5.2V21m4-3.2a2 2 0 0 1 4 0V21m11.794 0v-3.3a2 2 0 1 0-4 0V21m0-3.3v-2M31.306 21h0a1.847 1.847 0 0 1-1.842-1.841v-1.197a1.847 1.847 0 0 1 1.842-1.842h0a1.847 1.847 0 0 1 1.841 1.841v1.197A1.847 1.847 0 0 1 31.306 21Z"/>`),
		g.Group(children),
	)
}