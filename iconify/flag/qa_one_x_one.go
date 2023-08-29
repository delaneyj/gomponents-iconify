package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#8d1b3d" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 0v512h113l104.2-28.4L113 455l104.2-28.4L113 398.2l104.2-28.4L113 341.3L217.2 313L113 284.4L217.2 256L113 227.6L217.2 199L113 170.7l104.2-28.5L113 113.8l104.2-28.5L113 57l104.2-28.4L113 0H0z"/>`),
		g.Group(children),
	)
}