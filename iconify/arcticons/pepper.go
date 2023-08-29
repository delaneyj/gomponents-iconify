package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pepper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.644 28.856C38.644 36.943 32.088 43.5 24 43.5S9.356 36.944 9.356 28.856c0-5.642 5.369-6.931 5.369-6.931c-.968 2.803-1.828 11.5 5.51 11.5c2.35 0 3.492-1.332 3.492-3.5c0-3.652-7.256-7.612-7.256-14.459C16.47 8.483 24 4.5 24 4.5c-1.2 7.638 14.644 11.557 14.644 24.356Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.499 26.952c1.577 2.088 3.656 4.878 3.656 7.392c0 3.682-2.633 6.26-6.67 6.26c-5.033 0-6.956-2.644-6.956-2.644"/>`),
		g.Group(children),
	)
}