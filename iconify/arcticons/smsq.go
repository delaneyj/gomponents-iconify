package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smsq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.877 36.081c-1.526 1.494-.879 5.216.03 8.56l2.395-1.582c.03-1.587-.013-3.001.771-3.726h0a21.498 21.498 0 1 0-14.27 6.153a4.103 4.103 0 0 1 2.86 1.58l2.785-.705c-1.859-2.925-3.205-5.576-5.816-5.433a16.937 16.937 0 1 1 11.245-4.847Z"/><circle cx="16.959" cy="24" r="2.235" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.888" cy="24" r="2.235" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.041" cy="24" r="2.235" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}