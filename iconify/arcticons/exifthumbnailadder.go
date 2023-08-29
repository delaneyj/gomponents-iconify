package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exifthumbnailadder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.413 26.677l-4.39 5.61l-2.196-2.794l-4.579 5.713h5.056l6.798.037h5.73ZM8.7 7.98h0v4.2h0h-4.2h0v0a4.2 4.2 0 0 1 4.2-4.2Zm-4.2 6.96h4.2v4.2H4.5zm34.8-6.96h0a4.2 4.2 0 0 1 4.2 4.2v0h0h-4.2h0v-4.2h0Zm-27.84 0h4.2v4.2h-4.2zm6.96 0h4.2v4.2h-4.2zm6.96 0h4.2v4.2h-4.2zm6.96 0h4.2v4.2h-4.2zm0 27.84h4.2v4.2h-4.2zm11.16 0h0a4.2 4.2 0 0 1-4.2 4.2h0v-4.2h4.2Zm0-20.88v4.2h-4.2v-4.2zm0 6.96v4.2h-4.2v-4.2zm0 6.96v4.2h-4.2v-4.2z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 21.9h25.08v18.12h0H8.7a4.2 4.2 0 0 1-4.2-4.2V21.9h0Z"/>`),
		g.Group(children),
	)
}