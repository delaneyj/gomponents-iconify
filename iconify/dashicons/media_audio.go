package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zm0 4h3l-3-3v3zm1 7.26V8.09a.4.4 0 0 0-.12-.29a.3.3 0 0 0-.27-.1s-3.97.71-4.25.78C8.07 8.54 8 8.8 8 9v3.37c-.2-.09-.42-.07-.6-.07c-.38 0-.7.13-.96.39c-.26.27-.4.58-.4.96c0 .37.14.69.4.95c.26.27.58.4.96.4c.34 0 .7-.04.96-.26c.26-.23.64-.65.64-1.12V10.3l3-.6V12c-.67-.2-1.17.04-1.44.31c-.26.26-.39.58-.39.95c0 .38.13.69.39.96c.27.26.71.39 1.08.39c.38 0 .7-.13.96-.39c.26-.27.4-.58.4-.96z"/>`),
		g.Group(children),
	)
}