package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1378 384l64 128H606l64-128h708zm256 512l64 128H350l64-128h1220zm-128-256l64 128H478l64-128h964zm30-512l512 1024v640H0v-640L512 128h1024zM591 256l-448 896h1762l-448-896H591zm1329 1408v-384H128v384h1792zM512 1408v128H256v-128h256z"/>`),
		g.Group(children),
	)
}