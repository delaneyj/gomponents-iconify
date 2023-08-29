package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 277q-27 0-45.5-18.5T107 213q0-24 16.5-42.5T164 150h3q9-20 27.5-31.5T235 107q28 0 48.5 18t24.5 46h1q22 0 38 15.5t16 37.5t-16 37.5t-38 15.5H171zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h384zm0 342V42H43v300h384z"/>`),
		g.Group(children),
	)
}