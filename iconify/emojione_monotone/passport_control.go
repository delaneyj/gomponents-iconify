package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.431 2 2 15.43 2 32c0 16.568 13.432 30 30 30c16.569 0 30-13.432 30-30C62 15.43 48.568 2 32 2m11.814 11.471v3.801H34.26l-.949-3.801h10.503m-9.554 4.751h9.555v4.751c0 4.764-9.555 4.764-9.555 0v-1.9h-3.801l3.801-2.851m-9.622.831v4.371l-10.32-2.186l10.32-2.185m-10.32 4.371l10.32 2.186v9.834l-10.32-2.186v-9.834m7.06 22.353l-4.752-8.552h5.701l1.9 3.421l6.652-11.973h14.403L30.89 43.084v-4.168l-3.811 6.861h-5.701m26.304 4.751H30.89v-5.044h16.792v5.044m0-6.723H32.57l15.111-15.131v15.131z"/>`),
		g.Group(children),
	)
}