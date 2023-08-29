package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funkwhale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 23.093a18.5 18.5 0 0 1-37 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.335 23.093a11.335 11.335 0 1 1-22.67 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.463 23.093a5.463 5.463 0 0 1-10.926 0m5.173-.134c-2.689 0-.718-6.244-3.445-7.422c-2.938-1.269-9.27-1.195-9.988-9.13A27.884 27.884 0 0 1 18.1 7.911c3 1.284 4.821 1.198 6.208 4.582c0 0 .53-2.953 4.994-4.187a35.44 35.44 0 0 1 8.02-1.287a6.703 6.703 0 0 1-1.947 5.684c-2.53 2.332-7.963 2.273-8.441 4.652c-.49 2.432-.536 5.604-3.224 5.604Z"/>`),
		g.Group(children),
	)
}