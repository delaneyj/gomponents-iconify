package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyMoviesTV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 896v832q0 40-15 75t-41 61t-61 41t-75 15H320q-40 0-75-15t-61-41t-41-61t-15-75v-507q0-37 1-67t2-59t1-60t-4-67q-2-21-6-38t-8-34t-10-32t-11-38L22 541l1738-434l124 497L713 896h1207zM681 508l-321 80l352 176l321-80l-352-176zm543 129l322-81l-352-175l-322 80l352 176zm-1046 4l61 241l282-70l-343-171zm1489-379l-282 71l342 171l-60-242zm125 762H256v704q0 26 19 45t45 19h1408q26 0 45-19t19-45v-704z"/>`),
		g.Group(children),
	)
}