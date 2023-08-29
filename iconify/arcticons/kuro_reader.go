package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KuroReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.57 8.742v7.167m3.864-.008l-2.945-3.576l2.945-3.552m2.019 2.39v2.95c0 .988.8 1.788 1.788 1.788h0c.987 0 1.788-.8 1.788-1.788v-2.95m0 2.819v1.792m2.019-2.823c0-.987.8-1.788 1.788-1.788h0m-1.788 0v4.738m5.594 0h0c-.987 0-1.788-.8-1.788-1.788v-1.162c0-.987.8-1.788 1.788-1.788h0c.988 0 1.788.8 1.788 1.788v1.162c0 .988-.8 1.788-1.788 1.788Z"/><rect width="27" height="13.901" x="10.5" y="5.5" rx=".998" ry=".998"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24 24.64c-6.106-4.35-13.5-.107-13.5-.107V42.5s7.394-4.399 13.5-.065V24.641Z"/><path d="M24 24.64c-6.106-4.35-13.5-.107-13.5-.107V42.5s7.394-4.399 13.5-.065V24.641Zm0 17.795c6.106-4.334 13.5.065 13.5.065V24.533s-7.394-4.243-13.5.108v17.794Z"/></g>`),
		g.Group(children),
	)
}