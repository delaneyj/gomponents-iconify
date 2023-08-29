package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M10.05 18a20.46 20.46 0 0 0 .62 4.93h6.48v-9.48h-6.57a20.55 20.55 0 0 0-.53 4.55Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><path fill="currentColor" d="M18.85 22.94h6.48A20.46 20.46 0 0 0 26 18a20.55 20.55 0 0 0-.52-4.55h-6.63Z" class="clr-i-solid--badged clr-i-solid-path-2--badged"/><path fill="currentColor" d="M33.12 12.81a7.44 7.44 0 0 1-1.9.58H31a6.77 6.77 0 0 1-2.07 0h-1.8a21.88 21.88 0 0 1 .53 4.61a22.2 22.2 0 0 1-.57 4.93h4a13.94 13.94 0 0 1-.83 1.81H26.6a21.8 21.8 0 0 1-3 6a13.86 13.86 0 0 1-3 .92a20.21 20.21 0 0 0 4.18-6.94h-5.92v7.15h-1.69v-7.16h-5.95a20.21 20.21 0 0 0 4.18 6.95a13.86 13.86 0 0 1-2.94-.9a21.8 21.8 0 0 1-3-6.05H5.78a13.94 13.94 0 0 1-.83-1.81h4a22.2 22.2 0 0 1-.58-4.9a21.88 21.88 0 0 1 .48-4.55H4.76a13.88 13.88 0 0 1 .76-1.81h3.81A22.26 22.26 0 0 1 12.61 5a13.86 13.86 0 0 1 2.87-.84a20.13 20.13 0 0 0-4.4 7.45h6.09V4h1.69v7.64h6.09v-.13a7.47 7.47 0 0 1-2.36-4.76a20.37 20.37 0 0 0-2-2.55a14.23 14.23 0 0 1 2.06.56a7.44 7.44 0 0 1 .57-1.86a16.06 16.06 0 1 0 9.93 9.93Z" class="clr-i-solid--badged clr-i-solid-path-3--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-4--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}