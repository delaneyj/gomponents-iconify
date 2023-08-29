package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.454 19.028h-7.01l6.62-6.63a2.935 2.935 0 0 0 .87-2.09a2.844 2.844 0 0 0-.87-2.05l-3.42-3.44a2.93 2.93 0 0 0-4.13.01L3.934 13.4a2.946 2.946 0 0 0 0 4.14l1.48 1.49h-1.86a.5.5 0 0 0 0 1h16.9a.5.5 0 0 0 0-1.002Zm-7.24-13.5a1.956 1.956 0 0 1 2.73 0l3.42 3.44a1.868 1.868 0 0 1 .57 1.35a1.93 1.93 0 0 1-.57 1.37l-5.64 5.64l-6.15-6.16Zm-1.19 13.5h-5.2l-2.18-2.2a1.931 1.931 0 0 1 0-2.72l2.23-2.23l6.15 6.15Z"/>`),
		g.Group(children),
	)
}