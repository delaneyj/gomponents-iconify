package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M12.863 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M10.764 13L8.5 10.574v4.852L10.764 13Zm2.1.682a1 1 0 0 0 0-1.364L8.23 7.355c-.62-.664-1.731-.226-1.731.682v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/><path d="M17.863 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M15.764 13L13.5 10.574v4.852L15.764 13Zm2.1.682a1 1 0 0 0 0-1.364L13.23 7.355c-.62-.664-1.731-.226-1.731.682v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18.5 7a1 1 0 0 1 1 1v10a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}