package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.9 18.6L10.3 5.1h22.2c.8-.1 1.5.5 1.5 1.3v1.2c0 .5-.2 1-.6 1.4l-9.5 9.6z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M33.5 31L4.1 1.6L2.6 3l2.1 2.1H3.5C2.7 5 2 5.6 2 6.4v1.2c0 .5.2 1 .6 1.4L14 20.5v10.1l8 3.4V22.4l10.1 10.1l1.4-1.5z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}