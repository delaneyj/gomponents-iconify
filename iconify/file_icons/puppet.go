package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puppet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334.744 332.82V178.754H213.746l-60.674-61.692V0H0v153.072h120.006l60.674 61.692v82.47l-60.673 61.694H0V512h153.072V394.937l61.092-62.118h120.58zM103.237 102.803H51.154V50.268h52.083v52.536zm-.433 358.928H51.588v-51.216h51.216v51.216z"/>`),
		g.Group(children),
	)
}