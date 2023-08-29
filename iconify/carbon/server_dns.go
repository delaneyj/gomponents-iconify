package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerDns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a8 8 0 1 0 8 8a8.01 8.01 0 0 0-8-8zm5.91 7h-2.438a15.246 15.246 0 0 0-.791-4.36A6.009 6.009 0 0 1 21.91 9zm-5.888 6.999h-.008c-.38-.12-1.309-1.821-1.479-4.999h2.93c-.17 3.176-1.094 4.877-1.443 4.999zM14.535 9c.17-3.176 1.094-4.877 1.443-4.999h.008c.38.12 1.309 1.821 1.479 4.999zM13.32 4.64A15.246 15.246 0 0 0 12.528 9H10.09a6.009 6.009 0 0 1 3.23-4.36zM10.09 11h2.437a15.246 15.246 0 0 0 .792 4.36A6.009 6.009 0 0 1 10.09 11zm8.59 4.36a15.246 15.246 0 0 0 .792-4.36h2.438a6.009 6.009 0 0 1-3.23 4.36zM28 30H4a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zM4 22v6h24v-6z"/><circle cx="7" cy="25" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}