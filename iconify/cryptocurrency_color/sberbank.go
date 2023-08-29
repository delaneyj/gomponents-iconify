package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sberbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#48B254" fill-rule="nonzero"/><path fill="#FFF" d="m22.681 7.368l.945.858l-11.932 6.812l-5.776-3.325l.54-1.073l5.236 2.977l10.987-6.25zM20.279 6l1.268.644l-9.853 5.632l-4.588-2.602l.782-.938l3.806 2.172L20.28 6zm4.184 3.111l.701.939l-13.47 7.697l-6.505-3.701l.297-1.18l6.208 3.54l12.769-7.295zm1.943 3.46c.396 1.109.594 2.27.594 3.486c0 1.216-.198 2.397-.594 3.54l-.27.725a11.142 11.142 0 0 1-2.348 3.486a10.85 10.85 0 0 1-3.51 2.307c-1.385.59-2.815.885-4.291.885c-1.494 0-2.925-.295-4.293-.885a11.341 11.341 0 0 1-3.482-2.307a10.568 10.568 0 0 1-2.348-3.486c-.57-1.35-.865-2.8-.864-4.265v-.724l6.694 3.782l14.118-8.046l.594 1.502z"/></g>`),
		g.Group(children),
	)
}