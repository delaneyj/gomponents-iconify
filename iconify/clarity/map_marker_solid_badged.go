package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29.77 13.49A7.47 7.47 0 0 1 24.38 11a6.58 6.58 0 1 1-1.61-3a7.42 7.42 0 0 1 .31-4.84A11.75 11.75 0 0 0 6.22 13.73c0 4.67 2.62 8.58 4.54 11.43l.35.52a99.61 99.61 0 0 0 6.14 8l.76.89l.76-.89a99.82 99.82 0 0 0 6.14-8l.35-.53c1.91-2.85 4.53-6.75 4.53-11.42c-.01-.08-.02-.16-.02-.24Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="18" cy="12.44" r="3.73" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-3--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}