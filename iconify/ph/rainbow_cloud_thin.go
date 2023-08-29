package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowCloudThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 116a44.16 44.16 0 0 0-42 30.88a27.22 27.22 0 0 0-6-.66c-15.44 0-28 13-28 28.89S136.56 204 152 204h48a44 44 0 0 0 0-88Zm0 80h-48c-11 0-20-9.37-20-20.89s9-20.89 20-20.89a19.13 19.13 0 0 1 7.29 1.43a4 4 0 0 0 5.44-2.9A36 36 0 1 1 200 196ZM20 160v16a4 4 0 0 1-8 0v-16a100 100 0 0 1 169.71-71.69a4 4 0 0 1-5.57 5.69A92 92 0 0 0 20 160Zm92-60a60.07 60.07 0 0 0-60 60v16a4 4 0 0 1-8 0v-16a68 68 0 0 1 108.24-54.82a4 4 0 1 1-4.74 6.44A59.57 59.57 0 0 0 112 100Zm11.31 29.79a4 4 0 0 1-4.81 3A28 28 0 0 0 84 160v16a4 4 0 0 1-8 0v-16a36 36 0 0 1 36-36a36.58 36.58 0 0 1 8.35 1a4 4 0 0 1 2.96 4.79Z"/>`),
		g.Group(children),
	)
}