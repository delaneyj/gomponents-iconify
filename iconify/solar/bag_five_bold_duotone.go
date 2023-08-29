package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagFiveBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.035 16.927c.462 2.309.693 3.463 1.522 4.143c.83.68 2.007.68 4.362.68h4.162c2.355 0 3.532 0 4.361-.68c.83-.68 1.06-1.834 1.523-4.143l.6-3c.664-3.32.996-4.98.096-6.079c-.9-1.098-2.594-1.098-5.98-1.098H9.32c-3.386 0-5.08 0-5.98 1.098c-.9 1.098-.568 2.758.096 6.079l.6 3Z" opacity=".5"/><circle cx="15" cy="9.75" r="1"/><circle cx="9" cy="9.75" r="1"/><path d="M9.75 5.75a2.25 2.25 0 0 1 4.5 0v1h.431c.377 0 .733 0 1.069.002V5.75a3.75 3.75 0 1 0-7.5 0v1.002c.336-.002.692-.002 1.069-.002h.431v-1Z"/></g>`),
		g.Group(children),
	)
}