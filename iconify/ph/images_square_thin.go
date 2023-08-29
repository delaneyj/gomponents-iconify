package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H80a12 12 0 0 0-12 12v20H48a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h128a12 12 0 0 0 12-12v-20h20a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12ZM76 48a4 4 0 0 1 4-4h128a4 4 0 0 1 4 4v79l-23.51-23.52a12 12 0 0 0-17 0L95 180H80a4 4 0 0 1-4-4Zm104 160a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V80a4 4 0 0 1 4-4h20v100a12 12 0 0 0 12 12h100Zm28-28H106.34l70.83-70.83a4 4 0 0 1 5.66 0L212 138.34V176a4 4 0 0 1-4 4Zm-88-72a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm0-32a12 12 0 1 1-12 12a12 12 0 0 1 12-12Z"/>`),
		g.Group(children),
	)
}