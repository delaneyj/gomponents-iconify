package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M10.05 18a20.46 20.46 0 0 0 .62 4.93h6.48v-9.48h-6.57a20.55 20.55 0 0 0-.53 4.55Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M18.85 13.45v9.48h6.48A20.46 20.46 0 0 0 26 18a20.55 20.55 0 0 0-.52-4.55Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm12.22 22.71H26.6a21.8 21.8 0 0 1-3 6a13.86 13.86 0 0 1-3 .92a20.21 20.21 0 0 0 4.18-6.94h-5.92v7.15h-1.69v-7.13h-5.95a20.21 20.21 0 0 0 4.18 6.95a13.86 13.86 0 0 1-2.94-.9a21.8 21.8 0 0 1-3-6.05H5.78a13.94 13.94 0 0 1-.83-1.81h4a22.2 22.2 0 0 1-.58-4.9a21.88 21.88 0 0 1 .48-4.55H4.76a13.88 13.88 0 0 1 .76-1.81h3.81A22.26 22.26 0 0 1 12.61 5a13.86 13.86 0 0 1 2.87-.84a20.13 20.13 0 0 0-4.4 7.45h6.09V4h1.69v7.64h6.09a20.13 20.13 0 0 0-4.39-7.44a13.89 13.89 0 0 1 2.87.8a22.26 22.26 0 0 1 3.27 6.59h3.77a13.89 13.89 0 0 1 .76 1.81h-4.06a21.88 21.88 0 0 1 .49 4.6a22.2 22.2 0 0 1-.57 4.93h4a13.94 13.94 0 0 1-.87 1.78Z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}