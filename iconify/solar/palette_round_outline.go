package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteRoundOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 6a4.75 4.75 0 0 1 9.5 0v.432l2.033-2.033a4.75 4.75 0 1 1 6.718 6.717l-2.134 2.134H18a4.75 4.75 0 1 1 0 9.5H6A4.75 4.75 0 0 1 1.25 18V6Zm8.214 15.25H18a3.25 3.25 0 0 0 0-6.5h-2.133l-5.978 5.978a4.77 4.77 0 0 1-.425.522Zm1.286-3.504l7.69-7.69a3.25 3.25 0 1 0-4.596-4.597L10.75 8.553v9.193ZM6 2.75A3.25 3.25 0 0 0 2.75 6v12a3.25 3.25 0 0 0 6.5 0V6A3.25 3.25 0 0 0 6 2.75Zm0 15a.25.25 0 1 0 0 .5a.25.25 0 0 0 0-.5ZM4.25 18a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}