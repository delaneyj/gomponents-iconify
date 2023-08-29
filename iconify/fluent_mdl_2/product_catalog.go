package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductCatalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1792q0 27 10 50t27 40t41 28t50 10h640l257 128H256q-53 0-99-20t-82-55t-55-81t-20-100V256q0-49 21-95t57-82t82-57t96-22h1408v681l-128-64V128H256q-23 0-46 11t-41 30t-29 41t-12 46v1316q29-17 61-26t67-10h512v128H256q-27 0-50 10t-40 27t-28 41t-10 50zm1920-777v762l-576 287l-576-287v-762l576-287l576 287zm-576-144l-369 184l369 184l369-184l-369-184zm-448 827l384 191v-539l-384-192v540zm896 0v-540l-384 192v539l384-191z"/>`),
		g.Group(children),
	)
}