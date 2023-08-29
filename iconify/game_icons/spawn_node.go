package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpawnNode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M136 16v30h240V16H136zm0 60v60h240V76H136zm0 90v90h240v-90H136zm0 120v120h240V286H136zm-30 150c-15 0-30 15-30 30v30h360v-30c0-15-15-30-30-30H106z"/>`),
		g.Group(children),
	)
}