package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1536H256V384h1536v1152zM384 512v486l352-352l448 447l192-191l288 287V512H384zm0 896h933L736 827l-352 351v230zm1280 0v-37l-288-288l-102 101l225 224h165zm-192-640q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zM2048 0v2048H0V0h2048zm-128 128H128v1792h1792V128z"/>`),
		g.Group(children),
	)
}