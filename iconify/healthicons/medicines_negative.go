package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicinesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsMedicinesNegative0)"><path d="M38.282 14.985a1 1 0 0 1-1.198.75l-6.616-1.52a1 1 0 0 1 .448-1.95l6.616 1.521a1 1 0 0 1 .75 1.199ZM29.943 36.69l-3.381-7.25l4.597-2.144a4 4 0 1 1 3.38 7.25l-4.596 2.144Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0Zm-6 14a8 8 0 1 1-16 0a8 8 0 0 1 16 0ZM16.778 9.245c-.647-1.532-2.813-1.686-3.682-.263L6.292 20.115c-.785 1.284.092 2.924 1.623 3.033l11.866.846c1.53.11 2.644-1.388 2.06-2.77L16.778 9.245ZM35.385 36.36a6 6 0 0 0-5.071-10.876l-10.876 5.072a6 6 0 1 0 5.071 10.876l10.876-5.072Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMedicinesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}