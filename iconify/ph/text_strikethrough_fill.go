package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStrikethroughFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM82.71 94.58C86 76.57 104.58 64 128 64c18.2 0 33.59 7.41 41.18 19.83a8 8 0 1 1-13.66 8.34C150.94 84.66 140.39 80 128 80c-15.3 0-27.73 7.33-29.55 17.42a8 8 0 0 1-7.86 6.58a7.76 7.76 0 0 1-1.43-.13a8 8 0 0 1-6.45-9.29ZM192 136h-23.71a28.45 28.45 0 0 1 7.71 20c0 20.19-21.08 36-48 36c-23.89 0-43.83-12.78-47.43-30.4a8 8 0 1 1 15.67-3.2c2 9.87 16 17.6 31.76 17.6c17.35 0 32-9.16 32-20c0-9.14-6.76-14.43-25.72-20H64a8 8 0 0 1 0-16h128a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}