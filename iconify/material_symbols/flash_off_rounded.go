package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.075 13.225L7 4.15V4q0-.825.588-1.413T9 2h5.85q.8 0 1.288.625T16.425 4L15 9h1.125q.9 0 1.338.8t-.088 1.55l-1.3 1.875Zm3.7 9.4L13.75 16.6l-2.375 3.425q-.15.225-.388.3t-.462 0q-.225-.075-.375-.262T10 19.6V14H9q-.825 0-1.413-.588T7 12V9.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}