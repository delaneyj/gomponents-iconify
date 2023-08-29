package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PureNatural(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPureNatural0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M11 41c4.167-1.79 8.092-2.757 11.775-2.898c3.682-.142 8.424.49 14.225 1.898"/><path stroke-linecap="round" d="M23.045 44c-.767-9.44-.449-16.773.955-22"/><path fill="#fff" fill-rule="evenodd" stroke-linejoin="round" d="M24 23.176c1.59-6.018 4.39-9.607 8.4-10.766c4.009-1.16 7.876-.629 11.6 1.594c.019 4.892-2.105 8.476-6.37 10.751c-4.266 2.276-8.81 1.75-13.63-1.579Zm-.208-1.062c.833-7.187-.837-12.248-5.01-15.185C14.61 3.992 9.776 3.336 4.284 4.96c-.959 6.63.508 11.643 4.399 15.039c3.892 3.396 8.928 4.101 15.109 2.114Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPureNatural0)"/>`),
		g.Group(children),
	)
}