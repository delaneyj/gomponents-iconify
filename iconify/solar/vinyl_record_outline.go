package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.54 5.46A9.25 9.25 0 1 0 5.46 18.54A9.25 9.25 0 0 0 18.54 5.46ZM4.4 4.398C8.597.2 15.403.2 19.6 4.399c4.198 4.198 4.198 11.004 0 15.202s-11.003 4.199-15.2 0C.2 15.403.2 8.597 4.399 4.4Zm3.535 2.474a.75.75 0 0 1 0 1.061a5.75 5.75 0 0 0 0 8.132a.75.75 0 1 1-1.06 1.06a7.25 7.25 0 0 1 0-10.253a.75.75 0 0 1 1.06 0Zm8.132 0a.75.75 0 0 1 1.06 0a7.25 7.25 0 0 1 0 10.253a.75.75 0 1 1-1.06-1.06a5.75 5.75 0 0 0 0-8.132a.75.75 0 0 1 0-1.06ZM12 9.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}