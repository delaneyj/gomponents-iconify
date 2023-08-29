package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFrown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942Zm0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058Z"/><path fill="currentColor" d="M17.206 16.481a6.033 6.033 0 0 0-10.412 0a.5.5 0 0 0 .863.5a5.033 5.033 0 0 1 8.685 0a.5.5 0 0 0 .864-.5Z"/><circle cx="9" cy="9.011" r="1.25" fill="currentColor"/><circle cx="15" cy="9.011" r="1.25" fill="currentColor"/>`),
		g.Group(children),
	)
}