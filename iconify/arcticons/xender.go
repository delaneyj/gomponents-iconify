package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.955 23.993C12.307 31.497 8.845 36.9 7.487 39.663c-.3.611-1.03 3.792 3.184 3.792s9.084-4.87 13.329-4.87c4.245 0 9.115 4.87 13.329 4.87s3.485-3.18 3.184-3.792c-1.358-2.763-4.82-8.166-14.468-15.67c-8.99-6.993-9.74-8.07-9.74-10.036c0-1.92 3.231-2.529 9.552-2.529s10.535 2.388 10.535 2.388"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.045 23.918c9.648-7.504 13.11-12.908 14.468-15.67c.3-.612 1.03-3.793-3.184-3.793S28.245 9.325 24 9.325c-4.245 0-9.115-4.87-13.329-4.87S7.186 7.636 7.487 8.248c1.358 2.762 4.82 8.166 14.468 15.67c8.99 6.992 9.74 8.069 9.74 10.036c0 1.92-3.231 2.528-9.552 2.528s-10.535-2.388-10.535-2.388"/>`),
		g.Group(children),
	)
}