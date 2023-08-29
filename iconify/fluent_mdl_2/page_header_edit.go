package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageHeaderEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128v1792h640q-10 32-19 64t-17 64H0V0h1792v741q-35 5-66 16t-62 30V128H128zm1921 966q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q29-29 64-43t77-14q42 0 78 15t64 41t42 63t16 79zm-128 0q0-32-20-51t-52-19q-14 0-27 4t-23 15l-692 692l-34 135l135-34l692-691q21-21 21-51zm-385-198H256V258h1280v638zM384 384v384h1024V384H384zm0 1152h514l-25 25q-12 12-27 26q-5 20-9 39t-10 38H256v-640h1154l-129 128H384v384z"/>`),
		g.Group(children),
	)
}