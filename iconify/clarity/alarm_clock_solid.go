package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmClockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31.47 3.84a5.78 5.78 0 0 0-7.37-.63a16.08 16.08 0 0 1 8.2 7.65a5.73 5.73 0 0 0-.83-7.02Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M11.42 3.43a5.77 5.77 0 0 0-7.64.41a5.72 5.72 0 0 0-.38 7.64a16.08 16.08 0 0 1 8.02-8.05Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M18 4a14 14 0 0 0-9.89 23.88l-2.55 2.55A1 1 0 1 0 7 31.84l2.66-2.66a13.9 13.9 0 0 0 16.88-.08l2.74 2.74a1 1 0 0 0 1.41-1.41L28 27.78A14 14 0 0 0 18 4Zm7.47 17.43a1 1 0 0 1-1.33.47L17 18.44V9.69a1 1 0 0 1 2 0v7.5l6 2.91a1 1 0 0 1 .49 1.33Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}