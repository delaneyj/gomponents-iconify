package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dotenv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0H0zm74.178 463.246H32.49V421.56h41.687v41.686zm37.772-.285v-136.64h100.743l.208 24.54h-74.18v29.024h68.636l.196 23.085H138.72v35.008h74.918l.212 24.983h-101.9zm232.751 0h-28.785l-53.527-86.757v86.757h-25.765v-136.64h27.87l53.942 90.056V326.32h26.265v136.64zm92.732 0h-31.27l-48.02-136.64h31.304l33.411 95.734l31.754-95.734h31.557l-48.736 136.64z"/>`),
		g.Group(children),
	)
}