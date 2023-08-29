package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.976 3.433l3.646.002l1.384 1.956l1.374-1.956l3.643-.001l-5 8.406l-5.047-8.407z"/><path fill="currentColor" d="M14.6 2.43a1 1 0 0 0-.819.425L12 5.39l-1.791-2.537a1 1 0 0 0-.817-.423H6.38l3.55 5.92l2.1 3.5l2.07-3.5l3.52-5.92H14.6z"/><path fill="currentColor" d="m22.001 2.438l-4.384-.003L14.1 8.35l-2.07 3.5l-2.1-3.5l-3.546-5.913l-4.39.024a1 1 0 0 0-.857 1.506l10.02 17.105a1 1 0 0 0 1.727-.002l9.98-17.128a1 1 0 0 0-.863-1.504z" opacity=".5"/>`),
		g.Group(children),
	)
}