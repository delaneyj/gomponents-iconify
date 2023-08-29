package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBrokenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M33 7.64c-1.34-2.75-5.09-5-9.69-3.69a9.87 9.87 0 0 0-6 4.84a18.9 18.9 0 0 0-2.23 5.33l5.28 2.34l-4.6 4.37l3.49 4.1l1.52-1.3L18.54 21l5.4-5.13L17.58 13a16.23 16.23 0 0 1 2.17-4.1a7.68 7.68 0 0 1 4.11-3c3.34-.89 6.34.6 7.34 2.65c1.55 3.18.85 6.72-2.14 10.81A57.16 57.16 0 0 1 18 30.16A57.16 57.16 0 0 1 6.94 19.33c-3-4.1-3.69-7.64-2.14-10.81a5.9 5.9 0 0 1 5.33-2.93a7.31 7.31 0 0 1 2 .29a7.7 7.7 0 0 1 3.38 2l.15-.3a10.66 10.66 0 0 1 1-1.41a9.64 9.64 0 0 0-3.94-2.22C8.2 2.66 4.34 4.89 3 7.64c-1.88 3.85-1.1 8.18 2.32 12.87C8 24.18 11.83 27.9 17.39 32.22a1 1 0 0 0 1.23 0c5.55-4.31 9.39-8 12.07-11.71c3.41-4.69 4.19-9.02 2.31-12.87Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}