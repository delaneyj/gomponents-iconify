package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aggregator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.245 26.002c.087-.664.134-1.332.14-2.002a16.527 16.527 0 0 0-.14-2.002l4.334-3.393a1 1 0 0 0 .25-1.311l-4.104-7.117a1.001 1.001 0 0 0-1.251-.44l-5.115 2.061a15.693 15.693 0 0 0-3.463-2.002l-.771-5.435a1 1 0 0 0-1.001-.86h-8.228a1.001 1.001 0 0 0-1.001.86l-.77 5.435a15.379 15.379 0 0 0-3.464 2.002L9.526 9.736a1.001 1.001 0 0 0-1.252.44L4.17 17.294a1.001 1.001 0 0 0 .25 1.312l4.325 3.393A16.65 16.65 0 0 0 8.605 24c.006.67.053 1.338.14 2.002L4.42 29.395a1 1 0 0 0-.25 1.311l4.103 7.117a1 1 0 0 0 1.252.44l5.115-2.061a15.69 15.69 0 0 0 3.463 2.002l.77 5.435a1 1 0 0 0 1.002.86h8.208a.999.999 0 0 0 1-.86l.772-5.435a15.379 15.379 0 0 0 3.463-2.002l5.115 2.062a1 1 0 0 0 1.251-.44l4.104-7.118a1.001 1.001 0 0 0-.25-1.31l-4.294-3.394Z"/><circle cx="24" cy="24" r="12.012" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.495" cy="29.505" r="4.004" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.66 20.168A11.01 11.01 0 0 1 27.832 35.34"/>`),
		g.Group(children),
	)
}