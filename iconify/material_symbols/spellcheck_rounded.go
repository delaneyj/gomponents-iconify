package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpellcheckRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.1 19.2l4.95-4.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-4.925 4.925q-.3.3-.675.45t-.75.15q-.375 0-.75-.15t-.675-.45L10.55 18.45q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.15 2.15Zm-7.75-6.5l-.925 2.625q-.125.3-.375.488T4.475 16q-.55 0-.85-.45t-.1-.95l4.05-10.875Q7.7 3.4 7.988 3.2t.637-.2h.8q.35 0 .638.2t.412.525l4.05 10.85q.2.525-.113.975t-.887.45q-.35 0-.625-.2t-.4-.525l-.9-2.575H6.35Zm.7-1.9h3.9l-1.9-5.4h-.1l-1.9 5.4Z"/>`),
		g.Group(children),
	)
}