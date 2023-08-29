package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Post(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.522 15.105H43.5m-15.397 1.799l-3.077 4.224a3.973 3.973 0 1 0-2.739 6.858a4.44 4.44 0 0 0 3.179-.96c1.069-.838 5.918-7.49 8.056-11.92"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 15.105c-8.665 13.15-14.873 18.233-20.856 18.128s-9.564-4.123-9.564-9.227a9.202 9.202 0 0 1 14.942-7.23M9.3 23.067a12.115 12.115 0 0 1 3.726-8.017l-8.526.056m0-.001L9.299 23.1"/>`),
		g.Group(children),
	)
}