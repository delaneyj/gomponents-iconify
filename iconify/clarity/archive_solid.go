package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M19.41 20.6L18 22l-1.41-1.4L16 20H5v12a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V20H20ZM22 24a1 1 0 0 1-1 1h-6a1 1 0 0 1 0-2h6a1 1 0 0 1 1 1Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M30.5 12h-3.84v.13a3 3 0 0 1-.88 2.12L22 18h10v-4.5a1.5 1.5 0 0 0-1.5-1.5Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M10.2 14.25a3 3 0 0 1-.88-2.12s0-.09 0-.13H5.5A1.5 1.5 0 0 0 4 13.5V18h10Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="currentColor" d="m18 19.18l6.38-6.35A1 1 0 1 0 23 11.41l-4 3.95V3a1 1 0 1 0-2 0v12.4l-4-3.95a1 1 0 0 0-1.41 1.42Z" class="clr-i-solid clr-i-solid-path-4"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}