package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feWallet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWallet1" fill="currentColor"><path id="feWallet2" d="M20 9c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3Zm-2 0V6H4v12h14v-3h-2c-1.1 0-2-1.1-2-2v-1.968C14 9.9 14.9 9 16 9h2Zm-2 4h2v-2h-2v2Z"/></g></g>`),
		g.Group(children),
	)
}