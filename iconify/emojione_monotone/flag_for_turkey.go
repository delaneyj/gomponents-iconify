package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTurkey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-3.117 15.25c4.381 0 8.311 1.955 10.984 5.049a11.527 11.527 0 0 0-6.638-2.099c-6.444 0-11.668 5.283-11.668 11.8s5.224 11.8 11.668 11.8c2.466 0 4.753-.775 6.638-2.097c-2.674 3.092-6.604 5.047-10.984 5.047C20.828 46.75 14.3 40.147 14.3 32c0-8.146 6.528-14.75 14.583-14.75M49.7 36.254l-5.263-1.638l-3.246 4.268l.012-5.267L35.934 32l5.27-1.617l-.012-5.266l3.246 4.267l5.263-1.638L46.436 32l3.264 4.254"/>`),
		g.Group(children),
	)
}