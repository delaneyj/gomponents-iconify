package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraLensLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.858 19.71L12 16H5.07a8.018 8.018 0 0 0 4.788 3.71ZM4.252 14h4.284L5.07 7.999A7.963 7.963 0 0 0 4 12c0 .69.088 1.36.252 2Zm2.143-7.708L8.535 10L12 4a7.974 7.974 0 0 0-5.605 2.292Zm7.747-2.002L12 8h6.93a8.018 8.018 0 0 0-4.788-3.71ZM19.748 10h-4.284l3.465 6.001A7.964 7.964 0 0 0 20 12a8 8 0 0 0-.252-2Zm-2.143 7.708L15.465 14L12 20a7.974 7.974 0 0 0 5.605-2.292ZM12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm1.155-12h-2.31l-1.154 2l1.154 2h2.31l1.154-2l-1.154-2Z"/>`),
		g.Group(children),
	)
}