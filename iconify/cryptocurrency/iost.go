package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm8.5-21L16 6l-8.5 5v10l8.5 5l8.5-5zm-8.768 5.407l-1.031-.606l-.623.368l-1.476-.863l.628-.37l-3.33-1.959l6.271-3.688l5.917 3.438l-1.586.938l-4.327-2.53l-3.131 1.842l1.754 1.032l.897-.53l1.476.863l-.902.533l1.03.606l.736-.434l1.476.862l-.74.437l3.673 2.162l-6.27 3.687l-6.525-3.79l.027-1.843l6.495 3.787l3.13-1.841l-2.098-1.235l-.785.463l-1.476-.862z"/>`),
		g.Group(children),
	)
}