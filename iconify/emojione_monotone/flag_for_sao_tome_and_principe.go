package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSaoTomeAndPrincipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m40.802 31.061l-.935-2.995l-.935 2.995l-2.998.009l2.42 1.862l-.918 3.002l2.431-1.846l2.431 1.846l-.918-3.002l2.42-1.862zm11.8 0l-.935-2.995l-.936 2.995l-2.998.009l2.42 1.862l-.918 3.002l2.432-1.846l2.431 1.846l-.918-3.002l2.42-1.862z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-8.419 39.833L12.935 52.479a27.73 27.73 0 0 1-1.414-1.414L30.586 32L11.521 12.935a28.55 28.55 0 0 1 1.413-1.414L23.58 22.167h34.633C59.365 25.228 60 28.54 60 32s-.635 6.772-1.787 9.833H23.581"/>`),
		g.Group(children),
	)
}