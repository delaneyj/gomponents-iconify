package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudForEducation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 30h-2c0-1.654-1.346-3-3-3h-4c-1.654 0-3 1.346-3 3h-2c0-2.757 2.243-5 5-5h4c2.757 0 5 2.243 5 5zM15 16a1 1 0 0 0-1 1v6h2v-6a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M32 12H14v2h4v5c0 2.757 2.243 5 5 5s5-2.243 5-5v-5h4v-2Zm-9 10c-1.654 0-3-1.346-3-3v-1h6v1c0 1.654-1.346 3-3 3Zm3-6h-6v-2h6v2Z"/><path fill="currentColor" d="M25.798 10C24.87 5.441 20.83 2 16 2a9.977 9.977 0 0 0-9.822 8.124C2.655 10.754 0 13.849 0 17.5C0 21.636 3.365 25 7.5 25H12v-2H7.5A5.506 5.506 0 0 1 2 17.5a5.51 5.51 0 0 1 5.123-5.48l.836-.057l.09-.833A7.98 7.98 0 0 1 16 4c3.72 0 6.845 2.555 7.737 6h2.061Z"/>`),
		g.Group(children),
	)
}