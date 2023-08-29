package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplePodcastsLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 144a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm15.75 0h-31.5a20.17 20.17 0 0 0-16 7.82a19.67 19.67 0 0 0-3.58 17.05l12.18 48A20.17 20.17 0 0 0 124.44 232h7.12a20.17 20.17 0 0 0 19.64-15.13l12.18-48a19.67 19.67 0 0 0-3.58-17.05a20.17 20.17 0 0 0-16.05-7.82Zm38.49-2a8 8 0 0 0 5.75 9.74a8.12 8.12 0 0 0 2 .25a8 8 0 0 0 7.74-6a72 72 0 1 0-139.46 0a8 8 0 1 0 15.49-4a56 56 0 1 1 108.48 0ZM128 24a104 104 0 0 0-57.51 190.66a8 8 0 0 0 11.09-2.23a8 8 0 0 0-2.24-11.09a88 88 0 1 1 97.31 0a8 8 0 0 0 8.86 13.32A104 104 0 0 0 128 24Z"/>`),
		g.Group(children),
	)
}