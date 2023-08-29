package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.432 12.182l2.078 2.078h3.406v3.411l2.083 2.078v-7.568zm-21.765.511L0 19.308h1.745l.495-1.313h2.958l.547 1.313h1.781l-2.948-6.615zm5.614 0v6.609h1.646v-4.125l1.786 2.083h.25l1.786-2.073v4.12H15.4v-6.615h-1.443l-2.12 2.458l-2.125-2.458zm8.375 0v6.615h2.745c2.635 0 3.839-1.396 3.839-3.297c0-1.813-1.25-3.318-3.661-3.318zm1.651 1.213h1.052c1.563 0 2.177.948 2.177 2.094c0 .969-.5 2.094-2.156 2.094h-1.073zm-14.651.365l1.057 2.578H2.677zm22.854.411l-2.141 2.135v3h2.995l2.141-2.141H26.51z"/>`),
		g.Group(children),
	)
}