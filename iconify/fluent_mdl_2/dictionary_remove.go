package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DictionaryRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h766l128 128H256q-53 0-99-20t-82-55t-55-81t-20-100V256q0-49 21-95t58-82t82-57t95-22h1408v1219l-64 64l-64-64V128H256q-23 0-46 11t-41 30t-29 41t-12 46v1316q29-17 61-26t67-10h963l64 64l-64 64H256q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10zM1280 768H384V384h896v384zM512 512v128h640V512H512zm1469 797l-291 291l291 291l-90 90l-291-291l-291 291l-90-90l291-291l-291-291l90-90l291 291l291-291l90 90z"/>`),
		g.Group(children),
	)
}