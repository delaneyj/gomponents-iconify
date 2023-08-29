package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rviapay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 15.99a2.18 2.18 0 0 1 2.172-2.176h0m-2.172-.001v5.768m8.348-5.767l-2.173 5.767l-2.172-5.767m6.502-.001v5.768m6.501-2.177a2.18 2.18 0 0 1-2.172 2.177h0a2.18 2.18 0 0 1-2.172-2.176V15.99a2.18 2.18 0 0 1 2.172-2.176h0a2.18 2.18 0 0 1 2.172 2.176m0 3.591v-5.768m.011 15.172a3.522 3.522 0 0 1-3.512 3.512h0a3.522 3.522 0 0 1-3.511-3.512v-2.283a3.522 3.522 0 0 1 3.511-3.512h0a3.522 3.522 0 0 1 3.512 3.512m0 5.795V23.19m6.647 9.307l-3.687-9.307m7.023 0l-4.39 12.644c-.35.878-1.053 1.405-1.93 1.405h-.527M10.51 28.985a3.522 3.522 0 0 0 3.512 3.512h0a3.522 3.522 0 0 0 3.511-3.512v-2.283a3.522 3.522 0 0 0-3.511-3.512h0a3.522 3.522 0 0 0-3.511 3.512m0-3.512v14.049m10.494-25.933l1.52-.545"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}