package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsProtonCalendar0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.534 7.862A4.035 4.035 0 0 0 4.5 11.897v24.206a4.035 4.035 0 0 0 4.035 4.035h20.62v-9.862a4.483 4.483 0 0 1 4.483-4.483H43.5V11.897a4.035 4.035 0 0 0-4.035-4.035Z"/></defs><use href="#arcticonsProtonCalendar0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsProtonCalendar0" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 25.793V15.82a3.138 3.138 0 0 0-3.138-3.138H4.5m17.707 27.456V38.26a4.035 4.035 0 0 1 .964-2.617l7.105-8.338v.005m8.242 3.095l1.759-.957m0 0v7.035m-5.235-3.517a1.758 1.758 0 0 0 1.759-1.758h0a1.758 1.758 0 0 0-1.759-1.758"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.042 36.483a1.758 1.758 0 0 0 1.759-1.758h0a1.758 1.758 0 0 0-1.759-1.759m-2.663 2.924a2.95 2.95 0 0 0 2.17.594h.493m-2.664-6.449a2.95 2.95 0 0 1 2.173-.587l.492.002m-1.332 3.516h1.331"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.155 40.138h10.31a4.035 4.035 0 0 0 4.035-4.035v-10.31"/>`),
		g.Group(children),
	)
}