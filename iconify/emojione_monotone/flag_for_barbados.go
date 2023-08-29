package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBarbados(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m43.641 20.357l-1.057-1.059l-.006.008l.002-.008l-6.352 6.352h2.959c-.778 2.246-1.426 5.077-1.576 8.52h-4.156v-8.52h1.721L32 16.125l-3.175 9.525h1.72v8.52h-4.156c-.151-3.442-.798-6.273-1.577-8.52h2.954l-6.351-6.352l-2.114 2.111c.042.048 4.211 4.872 4.211 14.293v1.534h7.033v10.636h2.91V37.237h7.031v-1.534c0-9.421 4.172-14.245 4.213-14.293l-1.058-1.053"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM18.233 56.367V7.633C22.302 5.325 26.998 4 32 4s9.698 1.325 13.767 3.633v48.734C41.698 58.675 37.002 60 32 60s-9.698-1.325-13.767-3.633z"/>`),
		g.Group(children),
	)
}