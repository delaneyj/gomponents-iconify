package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyThreeHundredSixMNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBaby0306mNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18.848 31.034a3.5 3.5 0 1 0-5.696-4.068L11.699 29H9.5a3.5 3.5 0 1 0 0 7h4a3.5 3.5 0 0 0 2.848-1.466l2.5-3.5Zm20.626-.64H36.27l-.57-.263c-.428-.476-.935-1.116-1.486-1.814c-1.328-1.678-2.92-3.69-4.319-4.545A5.173 5.173 0 0 0 27.159 23H21.5l.586 13h4.729a5.332 5.332 0 0 0 2.273-.516a5.847 5.847 0 0 0 1.911-1.46l.57-.655c.224.249.442.598.67.961c.502.802 1.045 1.67 1.78 1.67h5.589c.67 0 1.178-.295 1.652-.821s.74-1.239.74-1.982c0-.744-.266-1.457-.74-1.982c-.474-.526-1.116-.822-1.786-.822ZM33.5 22a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBaby0306mNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}