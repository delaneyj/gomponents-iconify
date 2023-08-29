package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IffcoTokio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-30.239 7.053v8.325"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.174 18.068h0a2.798 2.798 0 0 1-2.81 2.81h0c-1.56 0-2.81-1.249-2.81-2.706v-2.81c0-1.56 1.25-2.81 2.81-2.705h0c1.561 0 2.706 1.248 2.706 2.705h0m-17.461 1.353h2.706m-2.706 4.163v-8.325h4.162m2.237 4.162h2.706m-2.706 4.163v-8.325h4.163m11.759 8.325h0c-1.562 0-2.707-1.249-2.707-2.706v-2.81c0-1.56 1.25-2.81 2.706-2.81h0a2.798 2.798 0 0 1 2.81 2.81v2.706a2.798 2.798 0 0 1-2.81 2.81Zm-4.156 6.244v8.325M9.5 27.122h5.412m-2.706 8.325v-8.325m6.816 8.325h0a2.687 2.687 0 0 1-2.706-2.705v-2.81c0-1.561 1.25-2.81 2.706-2.81h0a2.798 2.798 0 0 1 2.81 2.81v2.706a2.798 2.798 0 0 1-2.81 2.81Zm16.668 0h0a2.687 2.687 0 0 1-2.706-2.705v-2.81c0-1.561 1.25-2.81 2.706-2.81h0a2.798 2.798 0 0 1 2.81 2.81v2.706a2.798 2.798 0 0 1-2.81 2.81Zm-11.62-8.325v8.325m0-2.913l4.37-5.412m0 8.325l-3.33-4.162"/>`),
		g.Group(children),
	)
}