package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VendettaManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.015 14.527c.821.033 1.64.105 2.455.216l.79-1.827a13.02 13.02 0 0 1 7.417 2.775a35.442 35.442 0 0 1 3.787 15.417s-2.214 3.787-8.049 3.953a37.51 37.51 0 0 1-2.26-3.012m-10.176-.013a38.644 38.644 0 0 1-2.3 3.025c-5.835-.166-8.049-3.953-8.049-3.953a35.442 35.442 0 0 1 3.803-15.432a13.021 13.021 0 0 1 7.416-2.752l.79 1.827c.722-.1 1.447-.17 2.173-.21m0 18.246c-3.595-.263-6.247-1.647-8.98-3.031m20.43 0c-2.8 1.418-5.516 2.836-9.247 3.048"/><ellipse cx="19.556" cy="25.472" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.687" ry="1.977"/><ellipse cx="28.444" cy="25.472" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.687" ry="1.977"/><rect width="17.399" height="37" x="5.414" y="5.521" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><rect width="17.399" height="37" x="25.015" y="5.521" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.761 8.103h5"/><circle cx="14.113" cy="8.103" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}