package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v1.586l1-1L15.414 3L13 5.414v4.203l3.943-3.98l1.42 1.407l-3.949 3.987l4.169-.028L21 8.586L22.414 10l-1 1H23v2h-1.586l1 1L21 15.414l-2.411-2.411l-4.175.028l3.95 3.912l-1.407 1.421L13 14.446v4.14L15.414 21L14 22.414l-1-1V23h-2v-1.586l-1 1L8.586 21L11 18.586v-4.14l-3.957 3.918l-1.407-1.421l3.95-3.912l-4.175-.028L3 15.414L1.586 14l1-1H1v-2h1.586l-1-1L3 8.586l2.417 2.417l4.169.028l-3.95-3.988l1.421-1.407L11 9.617V5.414L8.586 3L10 1.586l1 1V1h2Z"/>`),
		g.Group(children),
	)
}