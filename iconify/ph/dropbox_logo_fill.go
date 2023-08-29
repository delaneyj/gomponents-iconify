package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m188 120l-60-40l55.56-37a8 8 0 0 1 8.88 0L238 73.34a8 8 0 0 1 0 13.32ZM72.44 43a8 8 0 0 0-8.88 0L18 73.34a8 8 0 0 0 0 13.32L68 120l60-40ZM238 153.34L188 120l-60 40l55.56 37a8 8 0 0 0 8.88 0L238 166.66a8 8 0 0 0 0-13.32Zm-220 0a8 8 0 0 0 0 13.32L63.56 197a8 8 0 0 0 8.88 0L128 160l-60-40Zm150.61 52.95l-38.37-25.58a4 4 0 0 0-4.44 0l-38.39 25.58a4 4 0 0 0 0 6.65L123.56 237a8 8 0 0 0 8.88 0l36.15-24.1a4 4 0 0 0 0-6.61Z"/>`),
		g.Group(children),
	)
}