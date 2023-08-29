package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm8.07 7.91L19.74 20H22a1 1 0 0 1 0 2h-2.75v2H22a1 1 0 0 1 0 2h-2.75v2.75a1.25 1.25 0 0 1-2.5 0V26H14a1 1 0 1 1 0-2h2.75v-2H14a1 1 0 1 1 0-2h2.26L9.93 9.91a1.25 1.25 0 1 1 2.12-1.33l5.95 9.5l5.95-9.49a1.25 1.25 0 1 1 2.12 1.33Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}