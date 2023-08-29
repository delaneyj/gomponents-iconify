package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M31 1a1 1 0 0 0-1 1v32a1 1 0 0 0 2 0V2a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M4 21v8a1 1 0 0 0 1 1h23V20H5a1 1 0 0 0-1 1Zm2 1h20v6H6Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M12 7v8a1 1 0 0 0 1 1h15V6H13a1 1 0 0 0-1 1Zm2 1h12v6H14Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}