package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnowCappedMountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.07 4.215a4.128 4.128 0 0 1 5.85 0l4.168 4.168l.012.002l6.9 6.9V29H2V15.285l11.07-11.07Zm4.183 4.316l.557-.279a8.908 8.908 0 0 1 2.847-.886l-2.444-2.444a3.128 3.128 0 0 0-4.436 0l-3.434 3.434l.428.211a6.975 6.975 0 0 0 3.079.718h.24c1.096 0 2.18-.254 3.157-.751l.006-.003ZM29 22.989l-1.003.448a19.407 19.407 0 0 1-7.91 1.695a19.32 19.32 0 0 1-6.35-1.072l-3.813-1.327A13.126 13.126 0 0 0 5.621 22c-.884 0-1.759.092-2.621.27V28h26v-5.011Z"/>`),
		g.Group(children),
	)
}