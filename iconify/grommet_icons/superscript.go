package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.74 3.63l.19 4.07h-.49a4.9 4.9 0 0 0-.38-1.54A2.57 2.57 0 0 0 17 5.07a3.68 3.68 0 0 0-1.73-.35h-2.44V18a3.56 3.56 0 0 0 .34 2a1.92 1.92 0 0 0 1.5.54h.6V21H7.92v-.5h.61a1.76 1.76 0 0 0 1.56-.67a3.88 3.88 0 0 0 .29-1.83V4.72H8.29a5.82 5.82 0 0 0-1.73.18a2.37 2.37 0 0 0-1.14.93a3.78 3.78 0 0 0-.56 1.87h-.48l.21-4.07ZM22.48 1h.33v5.62a1.13 1.13 0 0 0 .06.4a.41.41 0 0 0 .17.21a.63.63 0 0 0 .28.08h.4v.29h-2.96v-.27h.45a.62.62 0 0 0 .29-.1a.38.38 0 0 0 .15-.23a1.4 1.4 0 0 0 0-.37V2.77a4.45 4.45 0 0 0 0-.64c0-.15-.1-.23-.24-.23a.82.82 0 0 0-.28 0l-.3.13l-.18-.25Z"/>`),
		g.Group(children),
	)
}