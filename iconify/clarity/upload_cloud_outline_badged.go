package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadCloudOutlineBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M22.28 21.85a1 1 0 0 0 .72-1.71l-5-5l-5 5a1 1 0 0 0 1.41 1.41L17 19v12.25a1 1 0 1 0 2 0V19l2.57 2.57a1 1 0 0 0 .71.28Z" class="clr-i-outline--badged clr-i-outline-path-1--badged"/><path fill="currentColor" d="M30.92 13.44a7.13 7.13 0 0 1-2.63-.14v.23l-.08.72l.65.3A6 6 0 0 1 26.38 26H21v2h5.38a8 8 0 0 0 4.54-14.56Z" class="clr-i-outline--badged clr-i-outline-path-2--badged"/><path fill="currentColor" d="M3.6 18.38a7.71 7.71 0 0 1 7.4-7.64h.67l.23-.63a8.43 8.43 0 0 1 8-5.4a8.79 8.79 0 0 1 2.68.42a7.45 7.45 0 0 1 .5-1.94a10.79 10.79 0 0 0-3.18-.48a10.47 10.47 0 0 0-9.6 6.1a9.74 9.74 0 0 0-8.7 9.59a9.62 9.62 0 0 0 9.65 9.6H15v-2h-3.75a7.66 7.66 0 0 1-7.65-7.62Z" class="clr-i-outline--badged clr-i-outline-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-outline--badged clr-i-outline-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}