package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataPieTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.25 4.25A.75.75 0 0 1 11 5v8h8a.75.75 0 0 1 .743.648l.007.102c0 4.97-4.03 8.5-9 8.5a9 9 0 0 1-9-9c0-4.97 3.53-9 8.5-9ZM9.5 5.787l-.209.024c-3.69.47-6.041 3.622-6.041 7.439a7.5 7.5 0 0 0 7.5 7.5c3.817 0 6.968-2.352 7.44-6.041l.022-.209H10.25a.75.75 0 0 1-.743-.648L9.5 13.75V5.787Zm3.75-4.037a9 9 0 0 1 9 9a.75.75 0 0 1-.75.75h-8.25a.75.75 0 0 1-.75-.75V2.5a.75.75 0 0 1 .75-.75ZM14 3.287V10h6.712l-.023-.209a7.504 7.504 0 0 0-6.48-6.48L14 3.287Z"/>`),
		g.Group(children),
	)
}