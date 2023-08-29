package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.41 25.11a1 1 0 0 1-.54-1.85a6.21 6.21 0 0 0-.19-10.65a1 1 0 1 1 1-1.73A8.21 8.21 0 0 1 23.94 25a1 1 0 0 1-.53.11Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M18.34 3.87L9 12H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h5.83l9.51 8.3a1 1 0 0 0 1.66-.75V4.62a1 1 0 0 0-1.66-.75Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}