package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vala(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.384 24L8.868 2.14q-2.25.844-3.515 2.64q-1.25 1.797-1.25 4.47q0 .608.062 1q.078.374.156.609q.078.218.141.343q.078.125.078.22q-.828 0-1.468-.157q-.641-.172-1.079-.531q-.422-.359-.656-.953q-.22-.593-.22-1.469q0-1.062.453-2.093q.469-1.03 1.266-1.953q.812-.921 1.89-1.703q1.095-.781 2.329-1.344Q8.305.641 9.648.33Q11.008 0 12.352 0q.36 0 .656.015q.312.016.624.047l.282 19.687L20.648.155h2.234L13.992 24h-4.61Z"/>`),
		g.Group(children),
	)
}