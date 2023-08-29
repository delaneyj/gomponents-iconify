package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 4.25A16.49 16.49 0 0 0 5.4 31.4l.3.35h24.6l.3-.35A16.49 16.49 0 0 0 18 4.25Zm8.6 9.48l-5.92 5.81a3 3 0 1 1-1.41-1.42l5.91-5.81Zm-23 6.17H7v2H3.56c0-.39-.05-.77-.05-1.17s.02-.55.04-.83Zm4.88-10l2.46 2.46l-1.47 1.38L7 11.29a14.57 14.57 0 0 1 1.43-1.42ZM19 9.79h-2v-3.5h1c.37 0 .7 0 1 .05Zm13.49 10.95c0 .39 0 .79-.05 1.17h-3.52v-2h3.53c.02.27.04.55.04.83Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}