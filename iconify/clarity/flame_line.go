package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlameLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.3 16.66c-1.19-2.09-7.94-14.15-7.94-14.15a1 1 0 0 0-1.75 0l-6 10.64l-3-5.28a1 1 0 0 0-1.75 0S5.4 17.78 4.42 19.5A9.3 9.3 0 0 0 3 24.61C3 29.72 5.86 34 11.67 34h10.81C28.28 34 33 29 33 22.78a11.13 11.13 0 0 0-1.7-6.12ZM22.48 32H11.77C8.13 32 5 28.66 5 24.61a7.43 7.43 0 0 1 1.16-4.13c.73-1.29 4.05-7.21 5.65-10.07l3 5.28a1 1 0 0 0 .87.51a1 1 0 0 0 .87-.51L22.49 5c1.86 3.33 6.15 11 7.07 12.6A9.24 9.24 0 0 1 31 22.78c0 5.09-3.82 9.22-8.52 9.22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m25.75 21.73l-4.38-7.81a.8.8 0 0 0-1.4 0l-4.2 7.48l-1.59-2.49a.8.8 0 0 0-1.35 0l-3.46 5.44a4.35 4.35 0 0 0-.82 2.6a4.49 4.49 0 0 0 .5 2H11a3 3 0 0 1-.83-2a2.78 2.78 0 0 1 .56-1.73l2.8-4.38l1.66 2.6a.8.8 0 0 0 1.41-.12a7.82 7.82 0 0 1 .4-.8L20.67 16l3.69 6.57a4.83 4.83 0 0 1 .77 2.71A5 5 0 0 1 23.46 29h2.13a6.68 6.68 0 0 0 1.14-3.74a6.45 6.45 0 0 0-1-3.5Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}