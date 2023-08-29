package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageTableFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6.005A6.25 6.25 0 0 0 6.005 12H12V6.005ZM6 14.5v19h6v-19H6Zm8.5-2.5h19V6h-19v6ZM36 6.005V12h5.995A6.25 6.25 0 0 0 36 6.005Zm6 8.495h-6v19h6v-19ZM41.995 36H36v5.995A6.25 6.25 0 0 0 41.995 36ZM33.5 42v-6h-19v6h19ZM12 41.995V36H6.005A6.25 6.25 0 0 0 12 41.995Zm2.5-10.048l6.848-6.849a3.75 3.75 0 0 1 5.304 0l6.848 6.849V14.5h-19v17.447ZM26.5 19a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Zm5.018 14.5l-6.634-6.634a1.25 1.25 0 0 0-1.768 0L16.482 33.5h15.036Z"/>`),
		g.Group(children),
	)
}