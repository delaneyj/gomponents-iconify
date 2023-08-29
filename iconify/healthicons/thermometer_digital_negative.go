package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerDigitalNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsThermometerDigitalNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18.012 4.997l1.523 2.384l-7.686 4.909l-1.522-2.384l.008-.005a2.018 2.018 0 0 1-.115-.162l-.028-.044a2 2 0 0 1 .608-2.762l4.314-2.756a2 2 0 0 1 2.763.609l.028.044c.036.057.069.114.098.172l.009-.005Zm.663 17.979l-5.749-9l7.685-4.91l5.75 9c.613.962.77 1.595.977 2.428c.095.38.2.803.363 1.317c.071.225.133.492.209.823c.422 1.838 1.292 5.622 6.773 14.761a1.01 1.01 0 0 0-.137.073l-1.708 1.091a1 1 0 0 0-.128.097c-5.99-8.819-9.055-11.2-10.543-12.355a11.174 11.174 0 0 1-.658-.535a19.901 19.901 0 0 0-1.043-.883c-.668-.537-1.177-.946-1.79-1.907Zm18.44 18.344a178.2 178.2 0 0 1-1.401-2.231a1.38 1.38 0 0 1-.091.065l-1.708 1.09a.962.962 0 0 1-.103.058c.478.724.973 1.486 1.486 2.289l.596.933a1.108 1.108 0 1 0 1.867-1.192l-.646-1.012ZM24.057 20.107l-4.91-7.685l-2.561 1.636l4.909 7.685l2.562-1.636Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsThermometerDigitalNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}