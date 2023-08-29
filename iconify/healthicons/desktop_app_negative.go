package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopAppNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDesktopAppNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6 11.714C6 9.94 7.422 8.5 9.176 8.5h29.648C40.578 8.5 42 9.94 42 11.714V31c0 1.775-1.422 3.214-3.176 3.214H28.235v2.143h4.236V38.5H15.529v-2.143h4.236v-2.143H9.176C7.422 34.214 6 32.775 6 31V11.714Zm15.882 22.5v2.143h4.236v-2.143h-4.236ZM15.53 12.786h-2.117v2.143h-2.118v2.142h2.118v2.143h2.117v-2.143h2.118v-2.143H15.53v-2.142Zm-3.176 9.643c-.585 0-1.059.48-1.059 1.07v4.287c0 .591.474 1.071 1.059 1.071h4.235c.585 0 1.06-.48 1.06-1.071V23.5c0-.592-.475-1.072-1.06-1.072h-4.235Zm9.53-9.643c-.585 0-1.06.48-1.06 1.071v4.286c0 .592.475 1.071 1.06 1.071h4.235c.584 0 1.059-.48 1.059-1.071v-4.286c0-.592-.475-1.071-1.06-1.071h-4.235Zm0 9.643c-.585 0-1.06.48-1.06 1.07v4.287c0 .591.474 1.071 1.06 1.071h4.235c.584 0 1.058-.48 1.058-1.071V23.5c0-.592-.474-1.072-1.058-1.072h-4.236Zm9.529-9.643c-.585 0-1.06.48-1.06 1.071v4.286c0 .592.475 1.071 1.06 1.071h4.235c.585 0 1.059-.48 1.059-1.071v-4.286c0-.592-.474-1.071-1.059-1.071h-4.235Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDesktopAppNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}