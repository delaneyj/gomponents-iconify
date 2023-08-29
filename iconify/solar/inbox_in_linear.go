package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxInLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 2v8m0 0l3-3m-3 3L9 7"/><path d="M2 13h3.16c.905 0 1.358 0 1.756.183c.398.183.692.527 1.281 1.214l.606.706c.589.687.883 1.031 1.281 1.214c.398.183.85.183 1.756.183h.32c.905 0 1.358 0 1.756-.183c.398-.183.692-.527 1.281-1.214l.606-.706c.589-.687.883-1.031 1.281-1.214c.398-.183.85-.183 1.756-.183H22"/><path d="M17 2.127c1.625.16 2.72.521 3.535 1.338C22 4.929 22 7.286 22 12s0 7.071-1.465 8.536C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.464C2 19.07 2 16.714 2 12c0-4.714 0-7.07 1.464-8.535c.817-.817 1.91-1.178 3.536-1.338"/></g>`),
		g.Group(children),
	)
}