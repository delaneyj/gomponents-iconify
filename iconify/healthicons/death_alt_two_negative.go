package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeathAltTwoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDeathAlt2Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16.01 4l16 .013l4.99 13L31.978 44l-16-.017L11 16.992L16.01 4ZM23 17.997V26h2v-8.003h3v-2h-3V13h-2v2.997h-3v2h3Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDeathAlt2Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}