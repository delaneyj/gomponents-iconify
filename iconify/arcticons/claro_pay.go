package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClaroPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="4.202" height="5.568" x="28.632" y="24.049" rx="2.101" ry="2.101"/><path d="M25.142 26.15a2.1 2.1 0 0 1 2.101-2.101h0m-2.101 0v5.568m-8.097-8.405v8.405m6.107-2.101a2.1 2.1 0 0 1-2.1 2.1h0a2.1 2.1 0 0 1-2.102-2.1V26.15a2.1 2.1 0 0 1 2.101-2.101h0c1.16 0 2.102.94 2.102 2.1m-.001 3.468v-5.568m-8.084 2.749v.035a2.784 2.784 0 0 1-2.784 2.784h0A2.784 2.784 0 0 1 9.5 26.833v-2.837a2.784 2.784 0 0 1 2.784-2.784h0a2.784 2.784 0 0 1 2.784 2.784v.034m15.665-2.083v-3.564m3.168 4.941l3.152-3.152m-2.117 6.661H38.5"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M38.5 35.247v2.362a1.75 1.75 0 0 1-1.75 1.75h0c-.483 0-.92-.196-1.237-.512"/><path d="M38.5 32.36v2.887a1.75 1.75 0 0 1-1.75 1.75h0a1.75 1.75 0 0 1-1.75-1.75v-2.888m-1.767 2.888a1.75 1.75 0 0 1-1.75 1.75h0a1.75 1.75 0 0 1-1.75-1.75v-1.138c0-.966.784-1.75 1.75-1.75h0c.967 0 1.75.784 1.75 1.75m0 2.888v-4.638m-8.766 2.888c0 .966.783 1.75 1.75 1.75h0a1.75 1.75 0 0 0 1.75-1.75v-1.138a1.75 1.75 0 0 0-1.75-1.75h0a1.75 1.75 0 0 0-1.75 1.75m0-1.75v7"/></g>`),
		g.Group(children),
	)
}