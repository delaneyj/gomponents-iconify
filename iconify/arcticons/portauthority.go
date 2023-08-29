package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portauthority(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 8.5h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h11.948v-2.916h-2.252v-3.04h-5.74V14.456h27.09v19.09h-5.742v3.04h-2.253V39.5H41.5a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2Zm-23.052 31h11.103M24 18.841v-4.386m-2.222 4.386v-4.386m4.444 4.386v-4.386m4.437 4.386v-4.386m-13.342 4.386v-4.386m11.127 4.386v-4.386m-8.888 4.386v-4.386"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.112 14.455v4.386h17.776v-4.386"/>`),
		g.Group(children),
	)
}