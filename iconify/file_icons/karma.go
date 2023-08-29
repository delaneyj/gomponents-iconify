package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Karma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m208.46 452.658l-91.938-108.514l-57.504 54.172L0 180.665l91.95 108.517l57.492-54.197l59.018 217.673zm303.54 0l-139.196-205.66L500.162 59.342H376.005l-76.042 112.35V59.343H186.075v116.524l74.875 276.792h39.013v-98.332l11.051-16.283l76.829 114.615H512z"/>`),
		g.Group(children),
	)
}