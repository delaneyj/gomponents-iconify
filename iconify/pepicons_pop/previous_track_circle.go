package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousTrackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M13.137 12.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V8.037c0-.908-1.112-1.346-1.731-.682l-4.632 4.963Z"/><path fill-rule="evenodd" d="m15.236 13l2.264-2.426v4.852L15.236 13Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path d="M8.137 12.318a1 1 0 0 0 0 1.364l4.632 4.963c.62.664 1.731.226 1.731-.682V8.037c0-.908-1.112-1.346-1.731-.682l-4.632 4.963Z"/><path fill-rule="evenodd" d="m10.236 13l2.264-2.426v4.852L10.236 13Zm-2.1.682a1 1 0 0 1 0-1.364l4.633-4.963c.62-.664 1.731-.226 1.731.682v9.926c0 .908-1.112 1.346-1.731.682l-4.632-4.963Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.5 7a1 1 0 0 0-1 1v10a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}