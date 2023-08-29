package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#DB0A28" d="M20.99 59.62s-1.58 1.66-1.5 3.93c.09 2.87 1.92 4.12 3.83 6.99c1.92 2.87 34.64 49.1 36.06 50.58c1.63 1.7 2.28 3.07 4.7 3.02c2.4-.04 3.64-1.71 4.5-2.76c.86-1.05 36.85-51.51 39.05-54.67c1.37-1.97 1.15-3.55 1.05-4.41c-.1-.86-1.53-2.68-1.53-2.68H20.99z"/><path fill="#FF222D" d="m20.93 59.68l43.21 58.38l43.01-58.44S70.08 7.66 69.12 6.19c-.95-1.47-2.6-3.3-4.86-3.3c-2.17 0-3.9 1.56-5.47 3.73c-1 1.42-37.86 53.06-37.86 53.06z"/>`),
		g.Group(children),
	)
}