package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingerRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 255q-106 0-199 40T662 405T552 568t-41 200v768h708l128 128h-67q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100H256v-128h129V768q0-88 23-170t64-153t100-129t129-100t153-65t170-23q88 0 170 23t153 64t129 100t100 130t65 153t23 170v579l-127-126V768q0-106-40-199t-110-163t-163-110t-200-41zM904 1664q0 25 9 46t26 38t38 26t47 10q25 0 46-9t38-26t26-38t10-47H904zm915 64l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91l-226 226z"/>`),
		g.Group(children),
	)
}