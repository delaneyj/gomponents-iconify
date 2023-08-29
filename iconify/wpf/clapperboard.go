package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clapperboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m24.219.063l-4.844 1.25l2.438 1.437l-3.844 1L15.5 2.344l-4.813 1.25l2.438 1.437l-3.875 1l-2.438-1.406l-4.843 1.25l3.656 2.156L24.969 2.97l-.75-2.906zM2 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm4 3l-3 3h5l3-3H6zm9 0l-3 3h5l3-3h-5zm9 0l-3 3h5v-3h-2zM0 15v8a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3v-8H0z"/>`),
		g.Group(children),
	)
}