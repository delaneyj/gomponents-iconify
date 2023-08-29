package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileThirteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 0v128h128L320 0zm-21.3 0H64v512h384V149.3H298.7V0zm-192 42.7H256v106.7H106.7V42.7zM256 405.3H106.7v-42.7H256v42.7zM405.3 320H106.7v-42.7h298.7V320zm0-128v42.7H106.7V192h298.6z"/>`),
		g.Group(children),
	)
}