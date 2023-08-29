package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDuck0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 8c-3.052 4.517-5.987 5.81-8 5c.543 1.833 4.443 4.375 6.927 5.838c1.07.63 1.383 2.065.638 3.059C10.202 23.717 8.315 26.289 8 27C.954 39.79 16.482 44.648 24 44c22.144-1.908 21.355-19.197 18-26c-8.052 13.994-20.481 5.915-20 3c.481-2.915 3.792-2.335 5-7C29.013 4.768 16.374.399 12 8Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 31c2.5 3.5 10 7 16 2"/><circle cx="17" cy="12" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDuck0)"/>`),
		g.Group(children),
	)
}