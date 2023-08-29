package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M49.932 22.56a4.725 2.593 0 0 0-3.274.76L1.383 48.166a4.725 2.593 0 0 0 0 3.668L46.658 76.68a4.725 2.593 0 0 0 6.684 0l45.275-24.846a4.725 2.593 0 0 0 0-3.668L53.342 23.32a4.725 2.593 0 0 0-3.41-.76zM50 28.82l5.768 3.166l-38.592 21.18L11.406 50L50 28.82zm13.043 7.159l4.043 2.218l-38.594 21.18l-4.043-2.219l38.594-21.18zm11.318 6.21L88.596 50l-10.313 5.66l-14.234-7.81l10.312-5.66zm-15.51 8.51l14.235 7.813L50 71.18l-14.234-7.81L58.852 50.7z" color="currentColor"/>`),
		g.Group(children),
	)
}