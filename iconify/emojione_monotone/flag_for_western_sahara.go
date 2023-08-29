package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForWesternSahara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M38.555 32c0-3.328 2.253-6.107 5.245-6.743a6.03 6.03 0 0 0-1.31-.14c-3.622 0-6.557 3.084-6.557 6.883s2.935 6.883 6.557 6.883c.447 0 .888-.046 1.31-.141c-2.992-.635-5.245-3.414-5.245-6.742"/><path fill="currentColor" d="m45.715 31.017l-.934-2.951l-.906 2.951H40.85l2.42 1.914l-.916 3.002l2.427-1.845l2.43 1.845l-.916-3.002l2.422-1.914z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.431 2 32c0 16.569 13.432 30 30 30s30-13.432 30-30C62 15.431 48.568 2 32 2zM4 32c0-7.361 2.859-14.063 7.52-19.066L30.586 32L11.52 51.066C6.859 46.063 4 39.361 4 32zm19.581 9.832L33.414 32l-9.833-9.833h34.632A27.86 27.86 0 0 1 60 32a27.86 27.86 0 0 1-1.787 9.833H23.581z"/>`),
		g.Group(children),
	)
}