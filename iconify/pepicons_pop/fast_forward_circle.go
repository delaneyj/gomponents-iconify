package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M14.363 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M12.264 13L10 10.574v4.852L12.264 13Zm2.1.682a1 1 0 0 0 0-1.364L9.73 7.355C9.111 6.69 8 7.129 8 8.037v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/><path d="M19.363 12.318a1 1 0 0 1 0 1.364l-4.632 4.963c-.62.664-1.731.226-1.731-.682V8.037c0-.908 1.112-1.346 1.731-.682l4.632 4.963Z"/><path fill-rule="evenodd" d="M17.264 13L15 10.574v4.852L17.264 13Zm2.1.682a1 1 0 0 0 0-1.364L14.73 7.355c-.619-.665-1.73-.226-1.73.682v9.926c0 .908 1.112 1.346 1.731.682l4.632-4.963Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}