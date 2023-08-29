package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForQatar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM4 32c0-11.168 6.574-20.826 16.054-25.317L28.066 8.4l-9.833 2.949l9.833 2.95l-9.833 2.95l9.833 2.95l-9.833 2.95l9.833 2.949l-9.833 2.95L28.066 32l-9.833 2.95l9.833 2.95l-9.833 2.949l9.833 2.95l-9.833 2.95l9.833 2.95l-9.833 2.95l9.833 2.949l-8.01 1.719C10.575 52.828 4 43.169 4 32z"/>`),
		g.Group(children),
	)
}