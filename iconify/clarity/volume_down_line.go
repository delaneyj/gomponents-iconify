package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.41 25.11a1 1 0 0 1-.54-1.85a6.21 6.21 0 0 0-.19-10.65a1 1 0 1 1 1-1.73A8.21 8.21 0 0 1 23.94 25a1 1 0 0 1-.53.11Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 32a2 2 0 0 1-1.42-.59L9.14 24H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h5.22l7.33-7.41A2 2 0 0 1 20 6v24a2 2 0 0 1-1.24 1.85A2 2 0 0 1 18 32ZM4 14v8h5.56a1 1 0 0 1 .71.28L18 30V6l-7.65 7.68a1 1 0 0 1-.71.3Zm14-8Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}