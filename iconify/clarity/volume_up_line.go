package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.41 25.25a1 1 0 0 1-.54-1.85a6.21 6.21 0 0 0-.19-10.65a1 1 0 1 1 1-1.73a8.21 8.21 0 0 1 .24 14.06a1 1 0 0 1-.51.17Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M25.62 31.18a1 1 0 0 1-.45-1.89A12.44 12.44 0 0 0 25 6.89a1 1 0 1 1 .87-1.8a14.44 14.44 0 0 1 .24 26a1 1 0 0 1-.49.09Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M18 32.06a2 2 0 0 1-1.42-.59L9.14 24H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h5.22l7.33-7.39A2 2 0 0 1 20 6v24a2 2 0 0 1-1.24 1.85a2 2 0 0 1-.76.21ZM4 14v8h5.56a1 1 0 0 1 .71.3L18 30.06V6l-7.65 7.7a1 1 0 0 1-.71.3Zm14-8Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}