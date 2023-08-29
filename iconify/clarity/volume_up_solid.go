package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M23.41 25.25a1 1 0 0 1-.54-1.85a6.21 6.21 0 0 0-.19-10.65a1 1 0 1 1 1-1.73a8.21 8.21 0 0 1 .24 14.06a1 1 0 0 1-.51.17Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M25.62 31.18a1 1 0 0 1-.45-1.89A12.44 12.44 0 0 0 25 6.89a1 1 0 1 1 .87-1.8a14.44 14.44 0 0 1 .24 26a1 1 0 0 1-.49.09Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="m18.33 4l-9.26 8h-6a1 1 0 0 0-1 1v9.92a1 1 0 0 0 1 1h5.81l9.46 8.24a1 1 0 0 0 1.66-.73V4.72A1 1 0 0 0 18.33 4Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}