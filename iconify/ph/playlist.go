package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M32 64a8 8 0 0 1 8-8h176a8 8 0 0 1 0 16H40a8 8 0 0 1-8-8Zm8 72h120a8 8 0 0 0 0-16H40a8 8 0 0 0 0 16Zm72 48H40a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm135.66-57.7a8 8 0 0 1-10 5.36L208 122.75V192a32.05 32.05 0 1 1-16-27.69V112a8 8 0 0 1 10.3-7.66l40 12a8 8 0 0 1 5.36 9.96ZM192 192a16 16 0 1 0-16 16a16 16 0 0 0 16-16Z"/>`),
		g.Group(children),
	)
}