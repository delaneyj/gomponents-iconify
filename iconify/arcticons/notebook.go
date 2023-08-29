package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.3 15.172H12.664c-.758 0-1.372.614-1.372 1.372V39.18c0 .758.614 1.372 1.371 1.372H35.3c.757 0 1.372-.614 1.372-1.372V16.544c0-.758-.615-1.372-1.372-1.372Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m36.54 34.843l2.08.219a1.372 1.372 0 0 0 1.507-1.221l2.365-22.512a1.372 1.372 0 0 0-1.22-1.508L18.758 7.455a1.372 1.372 0 0 0-1.508 1.221l-.675 6.424"/><path d="M17.084 10.39L6.728 11.48c-.753.08-1.3.754-1.22 1.508l2.369 22.511c.08.754.754 1.3 1.508 1.221l1.768-.186m5.381-15.293H31.43m-14.896 4.414H31.43M16.534 30.07H31.43m-14.896 4.413h6.896"/></g>`),
		g.Group(children),
	)
}