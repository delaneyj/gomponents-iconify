package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeDuplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 640q0 95 27 185t77 167t120 138t158 101q15 7 37 15t46 15t47 13t39 6q18 1 36 1t36 0h72q36 0 73-1v768h-768v-768h216q-92-62-163-147T960 946q-45 102-116 187t-164 147h216v768H128v-768h72q36 0 73 1h36q18 0 36-1q16 0 39-5t47-13t46-16t37-15q87-39 157-100t121-139t77-167t27-185V347L605 637l-90-90l445-445l445 445l-90 90l-291-290v293zm-768 768v512h512v-512H256zm1408 0h-512v512h512v-512z"/>`),
		g.Group(children),
	)
}