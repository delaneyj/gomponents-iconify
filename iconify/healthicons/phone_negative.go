package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPhoneNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM26.42 34.76c-5.66-2.9-10.3-7.52-13.18-13.18l4.4-4.4c.56-.56.72-1.34.5-2.04A22.72 22.72 0 0 1 17 8c0-1.1-.9-2-2-2H8c-1.1 0-2 .9-2 2c0 18.78 15.22 34 34 34c1.1 0 2-.9 2-2v-6.98c0-1.1-.9-2-2-2c-2.48 0-4.9-.4-7.14-1.14c-.7-.24-1.5-.06-2.04.48l-4.4 4.4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPhoneNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}