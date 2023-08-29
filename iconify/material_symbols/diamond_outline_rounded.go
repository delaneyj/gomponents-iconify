package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19.875q-.425 0-.825-.188t-.7-.537L2.825 10q-.375-.45-.45-1.05t.2-1.125L4.45 4.1q.275-.5.738-.8T6.224 3h11.55q.575 0 1.038.3t.737.8l1.875 3.725q.275.525.2 1.125t-.45 1.05l-7.65 9.15q-.3.35-.7.538t-.825.187ZM9.625 8h4.75l-1.5-3h-1.75l-1.5 3ZM11 16.675V10H5.45L11 16.675Zm2 0L18.55 10H13v6.675ZM16.6 8h2.65l-1.5-3H15.1l1.5 3ZM4.75 8H7.4l1.5-3H6.25l-1.5 3Z"/>`),
		g.Group(children),
	)
}