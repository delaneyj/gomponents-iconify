package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashscore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.424 38.246a14.474 14.474 0 0 0 5.265-.465l-1.766-5.509a8.554 8.554 0 0 1-2.918.225l-.58 5.749Zm-4.486-1.191a14.361 14.361 0 0 1-5.272-4.076l4.503-3.62a8.562 8.562 0 0 0 3.23 2.466l-2.461 5.23Zm-7.027-6.85a14.256 14.256 0 0 1-1.388-7.031l5.754.826a8.5 8.5 0 0 0 .768 3.542L5.91 30.205Zm16.292-20.113a14.347 14.347 0 0 0-3.39-.404c-6.691 0-12.31 4.59-13.877 10.794l5.571 1.542a8.538 8.538 0 0 1 9.621-6.458l2.075-5.474Zm-2.852 12.219l6.348-12.623H43.5L19.351 22.311zM27.347 24h5.778m-2.449 8.009a14.39 14.39 0 0 1-2.565 2.872l-3.874-4.291a8.586 8.586 0 0 0 1.633-1.79l4.806 3.209Z"/>`),
		g.Group(children),
	)
}