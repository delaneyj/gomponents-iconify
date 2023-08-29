package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoapps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.928 10.53c.542.5 1.056 1.014 1.542 1.542l5.846.836a22.102 22.102 0 0 1 2.184 5.267l-3.555 4.733s.064 1.456 0 2.184l3.555 4.733a21.883 21.883 0 0 1-2.184 5.268l-5.846.835s-1.007 1.05-1.542 1.542l-.835 5.846a22.102 22.102 0 0 1-5.268 2.184l-4.733-3.555c-.727.065-1.457.065-2.184 0L18.175 45.5a21.883 21.883 0 0 1-5.268-2.184l-.835-5.846a39.263 39.263 0 0 1-1.542-1.542l-5.846-.835A22.098 22.098 0 0 1 2.5 29.825l3.555-4.733s-.065-1.456 0-2.184L2.5 18.175a21.883 21.883 0 0 1 2.184-5.268l5.846-.835a40.04 40.04 0 0 1 1.542-1.542l.835-5.846A22.098 22.098 0 0 1 18.175 2.5l4.733 3.555a12.391 12.391 0 0 1 2.184 0L29.825 2.5a21.883 21.883 0 0 1 5.268 2.184l.835 5.846Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.566 14.339l-9.8 9.8h10.469l-9.523 9.522"/>`),
		g.Group(children),
	)
}