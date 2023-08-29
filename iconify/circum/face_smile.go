package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSmile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942Zm0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058Z"/><path fill="currentColor" d="M16.693 13.744a5.041 5.041 0 0 1-9.387 0c-.249-.59-1.111-.081-.863.505a6.026 6.026 0 0 0 11.114 0c.247-.586-.614-1.1-.864-.505Z"/><circle cx="9" cy="9.011" r="1.25" fill="currentColor"/><circle cx="15" cy="9.011" r="1.25" fill="currentColor"/>`),
		g.Group(children),
	)
}