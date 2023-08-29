package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.86 2.504a1.5 1.5 0 0 0-2.415 1.708L3.828 19.828a6 6 0 1 0 8.486 8.486L27.93 12.697a1.5 1.5 0 0 0 1.708-2.414L21.86 2.504Zm-1.061 3.182l5.657 5.657l-3.068 3.068l-10.476-.838l1.877-1.876l1.59 1.59a.75.75 0 0 0 1.061-1.06l-1.59-1.59l4.949-4.95Zm-9.192 9.193l1.59 1.59a.75.75 0 0 1-1.06 1.061l-1.591-1.59l1.06-1.061ZM7.364 19.12l1.591 1.591a.75.75 0 0 1-1.06 1.061l-1.592-1.591l1.061-1.06Z"/>`),
		g.Group(children),
	)
}