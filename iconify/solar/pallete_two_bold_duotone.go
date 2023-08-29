package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalleteTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.847 21.934C5.867 21.362 2 17.133 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10c0 5.157-3.283 4.733-6.086 4.37c-1.618-.209-3.075-.397-3.652.518c-.395.626.032 1.406.555 1.929a1.673 1.673 0 0 1 0 2.366c-.523.523-1.235.836-1.97.751Z" opacity=".5"/><path d="M11.085 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM6.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm11 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-3-4.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g>`),
		g.Group(children),
	)
}