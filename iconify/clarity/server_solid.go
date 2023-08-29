package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M26.5 2h-17C8.7 2 8 2.7 8 3.5V34h20V3.5c0-.8-.7-1.5-1.5-1.5zM18 30.5c-1.5 0-2.8-1.2-2.8-2.8S16.5 25 18 25s2.8 1.2 2.8 2.8s-1.3 2.7-2.8 2.7zm5-7.9H13V21h10v1.6zm1-11H12V10h12v1.6zm0-4H12V6h12v1.6z" class="clr-i-solid clr-i-solid-path-1"/><circle cx="18" cy="27.8" r="1.2" fill="currentColor" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}