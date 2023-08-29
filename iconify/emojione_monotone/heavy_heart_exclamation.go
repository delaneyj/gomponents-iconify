package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyHeartExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="52" r="10" fill="currentColor"/><path fill="currentColor" d="M57.411 11.809C52.925-2.703 34.297 1.155 32 8.862C29.705 1.155 11.077-2.703 6.59 11.809C4.638 18.123 7.723 24 13.038 27.715C20.152 32.687 28.542 32.707 32 42c3.458-9.293 11.849-9.313 18.963-14.285c5.316-3.715 8.399-9.592 6.448-15.906z"/>`),
		g.Group(children),
	)
}