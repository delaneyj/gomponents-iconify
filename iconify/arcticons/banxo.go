package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banxo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.709 27.988h12.813M26.157 42.5h14.807m-10.362-6.799h7.404M12.2 18.836h12.813M9.455 23.412h12.813m15.035-2.615l3.988 4.51l-2.746 2.681M38.61 5.5c-15.492 7.06-37.457 21.376-19.48 37l15.1-25.756"/><circle cx="35.178" cy="25.34" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}