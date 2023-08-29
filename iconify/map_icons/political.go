package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Political(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M12.333 22H15v16h4V24h4v14h4V24h4v14h4V22h2.666c.736 0 1.334-.736 1.334-1.399c0-.401-.221-.853-.555-1.071l.002-.052l-.021-.037l.134-.033L26 12.939v-1.892c2 1.059 3.951-.765 7 .332V7.2c-3.051-1.096-5 .727-7-.332V6.6c0-.331-.631-.6-1-.6s-1 .269-1 .6v6.339l-12.559 6.456l.271.021l.288.011v.002c-1 .218-.776.77-.776 1.171c-.001.664.373 1.4 1.109 1.4zM14 39l-3 3h28l-3-3zm22-16v15h.885l.391.553l2.666 2.499l.943.948H49V23H36zm6 12h-3v-3h3v3zm0-6h-3v-3h3v3zm4 6h-2v-3h2v3zm0-6h-2v-3h2v3zm-32.886 9H14V23H1v19h8.114l.942-.848l2.667-2.6l.391-.552zM6 35H4v-3h2v3zm0-6H4v-3h2v3zm5 6H8v-3h3v3zm0-6H8v-3h3v3z"/>`),
		g.Group(children),
	)
}