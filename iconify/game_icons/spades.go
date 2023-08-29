package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spades(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M458.915 307.705c0 62.63-54 91.32-91.34 91.34c-41.64 0-73.1-18.86-91.83-34.26c2.47 50.95 14.53 87.35 68.65 116h-176.79c54.12-28.65 66.18-65.05 68.65-116c-18.73 15.39-50.2 34.28-91.83 34.26c-37.29 0-91.34-28.71-91.34-91.34c0-114.47 80.64-83.32 202.91-276.49c122.28 193.17 202.92 162.03 202.92 276.49z"/>`),
		g.Group(children),
	)
}