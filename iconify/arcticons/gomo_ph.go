package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GomoPh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.858 31.329V21.198l3.798 10.142l4.518-12.046V31.34m-12.165 0h0c-1.037 0-1.839-1.31-1.839-2.924v-2.979c0-1.614.802-3.109 1.839-3.349h0c1.093-.252 2.024.974 2.024 2.752v3.28c0 1.778-.931 3.22-2.024 3.22Zm19.095 0h0c-2.235 0-3.926-1.892-3.926-4.226v-4.306c0-2.334 1.691-4.617 3.926-5.134h0c2.417-.559 4.526 1.135 4.526 3.827v4.965c0 2.692-2.11 4.874-4.527 4.874Zm-22.61-5.69c0-1.555-.716-2.652-1.56-2.456h0c-.807.186-1.434 1.49-1.434 2.92v2.637c0 1.43.627 2.589 1.433 2.589h0c.845 0 1.56-1.262 1.56-2.818l-1.56.12"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m36.278 30.162l-.279 1.245l.939.865l1.217-.381l.279-1.245l-.938-.865l-1.218.381zm2.631-11.541v1.304l1.129.651l1.128-.651v-1.304l-1.128-.652l-1.129.652z"/><path d="M39.694 15.728c2.4 0 3.806 1.237 3.806 3.224c0 2.117-3.284 9.018-5.26 9.018c-1.188 0-1.309-.873-1.309-1.818V18.54c0-.92.655-2.812 2.763-2.812Z"/></g>`),
		g.Group(children),
	)
}