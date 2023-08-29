package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2A11.79 11.79 0 0 0 6.22 13.73c0 4.67 2.62 8.58 4.54 11.43l.35.52a99.61 99.61 0 0 0 6.14 8l.76.89l.76-.89a99.82 99.82 0 0 0 6.14-8l.35-.53c1.91-2.85 4.53-6.75 4.53-11.42A11.79 11.79 0 0 0 18 2Zm0 17a6.56 6.56 0 1 1 6.56-6.56A6.56 6.56 0 0 1 18 19Z" class="clr-i-solid clr-i-solid-path-1"/><circle cx="18" cy="12.44" r="3.73" fill="currentColor" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}