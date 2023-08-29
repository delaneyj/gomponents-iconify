package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousTrackOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M10.137 9.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V5.037c0-.908-1.112-1.346-1.731-.682l-4.632 4.963Z"/><path fill-rule="evenodd" d="M12.236 10L14.5 7.574v4.852L12.236 10Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path d="M5.137 9.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V5.037c0-.908-1.112-1.346-1.731-.682L5.137 9.318Z"/><path fill-rule="evenodd" d="M7.236 10L9.5 7.574v4.852L7.236 10Zm-2.1.682a1 1 0 0 1 0-1.364L9.77 4.355c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.5 4a1 1 0 0 0-1 1v10a1 1 0 1 0 2 0V5a1 1 0 0 0-1-1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}