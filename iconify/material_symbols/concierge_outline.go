package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConciergeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22v-2h13v2H10Zm1-3q0-2.025 1.275-3.538T15.5 13.6v-.625q0-.425.288-.712t.712-.288q.425 0 .713.288t.287.712v.625q1.925.35 3.213 1.863T22 19H11Zm2.625-2.025h5.7q-.475-.675-1.213-1.088t-1.612-.412q-.9 0-1.65.413t-1.225 1.087Zm2.85 0ZM1 13V2h6v1.45l7-1.95L22 4v1q0 1.25-.875 2.125T19 8h-2v.6q0 .625-.363 1.137t-.937.738L8.95 13H1Zm2-2h2V4H3v7Zm4 0h1.6l5.8-2.125q.275-.1.438-.338T15 8h-1.775l-2.925.95l-.6-1.9L12.825 6H19q.225 0 .563-.163t.337-.387L13.95 3.6L7 5.5V11Z"/>`),
		g.Group(children),
	)
}