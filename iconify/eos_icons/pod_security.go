package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.993 4.012H6.988V2.998a.998.998 0 0 1 1-.998h4.007a.998.998 0 0 1 .998.998v1.014ZM12 15.994H2l2.5-5.496L7 5.003h5.993l1.308 2.871l1.309 2.872l-1.805.802L12 12.35v3.644ZM12 10a2 2 0 1 0-.586 1.414A1.994 1.994 0 0 0 12 10Zm.085 6.995H3.993V18a2.001 2.001 0 0 0 2.002 2h7.239a7.649 7.649 0 0 1-.737-1.432a7.191 7.191 0 0 1-.412-1.573ZM22 13v3a6.405 6.405 0 0 1-1.282 3.804A5.776 5.776 0 0 1 17.5 22a5.776 5.776 0 0 1-3.217-2.196A6.405 6.405 0 0 1 13 16v-3l2.25-1l2.25-1l2.25 1Zm-4.5-.905l-1.75.777l-1.75.778v2.85h3.5v4.465a4.788 4.788 0 0 0 2.348-1.678A5.803 5.803 0 0 0 21 16.495h-3.5v-4.4Z"/>`),
		g.Group(children),
	)
}