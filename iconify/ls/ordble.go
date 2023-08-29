package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ordble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M356 0c96 0 179 35 249 105c70 71 105 156 105 256c0 99-35 184-104 253S454 717 357 717c-102 0-188-35-256-106S0 456 0 359c0-66 15-126 47-180s75-98 130-130c5-3 11-7 17-9c0 0 58 35 57 94c-51-25-105 28-105 28s58-16 93 9c-85 50-130 153-103 253c32 121 156 193 277 160s192-156 159-276c-30-113-138-182-250-166c8-39 50-81 50-81s-44-6-79 57c-11-69-98-78-99-78C244 13 298 0 356 0zM222 238c23-23 54-29 69-14c14 15 8 45-16 68c-23 24-53 30-68 15c-14-14-9-45 15-69z"/>`),
		g.Group(children),
	)
}