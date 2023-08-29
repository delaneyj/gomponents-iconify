package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcecreamFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIcecreamFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m18.843 13.723l13.435 13.435L11.679 40.5l-6.045-6.045l13.209-20.732Z"/><path fill="#fff" d="M38.773 19.51c-1.06 1.061-1.76 1.07-2.812 1.431c-.033 2.796.006 5.61-3.542 6.304L18.603 13.429s-.009-1.37 1.396-2.085c.495-.232 1.499-.123 2.097-.025c1.406-3.025 3.535-3.536 5.903-2.558c1.135-1.718 3.94-2.45 6.028-1.775c2.089.675 2.072 2.073 3.462 2.755c1.39.683 2.804-.731 2.812-1.43c.008-.7-.114-2.036.828-1.794c.575.156 1.038.715 1.596 3.044c.31 1.308.13 3.242-.987 4.429c-.686.729-2.217 1.11-2.916 1.327c.066.19 1.012 3.133-.05 4.193Z"/><path d="M38.822 15.316c-2.813 1.43-4.901.756-6.974-1.316"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIcecreamFour0)"/>`),
		g.Group(children),
	)
}