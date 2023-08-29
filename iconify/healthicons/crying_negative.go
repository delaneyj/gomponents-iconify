package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CryingNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCryingNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 40c2.68 0 5.206-.659 7.424-1.823l.955 1.758A17.922 17.922 0 0 1 24 42c-9.941 0-18-8.059-18-18S14.059 6 24 6s18 8.059 18 18c0 .167-.002.334-.007.5h-2c.004-.166.007-.333.007-.5c0-8.837-7.163-16-16-16S8 15.163 8 24s7.163 16 16 16Zm16-8a4 4 0 0 1-8 0c0-3.5 4-7 4-7s4 3.5 4 7Zm-19.2 1.6c1.6-2.133 4.8-2.133 6.4 0a1 1 0 0 0 1.6-1.2c-2.4-3.2-7.2-3.2-9.6 0a1 1 0 0 0 1.6 1.2Zm-5.993-8.367c-.683-.438-1.183-1.145-1.064-1.883a.5.5 0 0 1 .668-.389c1.818.675 3.846.257 5.642-1.448a.5.5 0 0 1 .818.203c.231.689.024 1.618-.48 2.382a3.794 3.794 0 0 1-2.654 1.665c-1.198.177-2.216-.073-2.93-.53Zm17.947 0c.683-.438 1.184-1.145 1.065-1.883a.5.5 0 0 0-.668-.389c-1.818.675-3.846.257-5.642-1.448a.5.5 0 0 0-.818.203c-.232.689-.024 1.618.48 2.382a3.794 3.794 0 0 0 2.653 1.665c1.199.177 2.216-.073 2.93-.53Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCryingNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}