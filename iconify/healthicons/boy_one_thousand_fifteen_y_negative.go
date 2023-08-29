package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyOneThousandFifteenYNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBoy1015yNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm18 18.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.415 2H30v.034c2.25.304 4 2.635 4 5.466c0 2.831-1.75 5.162-4 5.466V40a2 2 0 0 1-3.988.22L25.21 33h-2.32l-.905 7.248A2 2 0 0 1 18 40V29.966c-2.25-.304-4-2.635-4-5.466c0-2.831 1.75-5.162 4-5.466V19h.085a1.497 1.497 0 0 1-.085-.5Zm2.331 5.18c.478-.987.662-2.53.662-4.68h2c0 2.158-.166 4.115-.862 5.551c-.361.746-.877 1.379-1.6 1.817c-.72.435-1.571.632-2.531.632v-2c.686 0 1.158-.14 1.495-.344c.334-.201.61-.51.836-.976ZM27.007 19c0 2.15.184 3.693.662 4.68c.226.465.502.775.836.976c.337.204.809.344 1.495.344v2c-.96 0-1.81-.197-2.53-.632c-.724-.438-1.24-1.071-1.601-1.817c-.696-1.436-.862-3.393-.862-5.551h2ZM24 15a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBoy1015yNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}