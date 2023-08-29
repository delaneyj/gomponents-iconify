package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHealthNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM27.408 6.8a5.537 5.537 0 0 0-5.737 0L19.72 7.995c-3.267 1.996-7.248 2.374-11.101 2.74c-.873.083-1.74.166-2.59.266c0 0-.024 1.845 0 2.689l.209 7.128c.179 6.133 2.906 11.948 7.585 16.174c2.448 2.211 5.13 4.037 8.158 5.477c1.622.77 3.546.7 5.101-.187c2.623-1.496 4.906-3.297 7.113-5.29c4.678-4.226 7.406-10.041 7.585-16.174l.209-7.151c.024-.832.001-2.666.001-2.666a70.828 70.828 0 0 0-3.162-.426c-3.27-.392-6.526-.781-9.383-2.528l-2.037-1.245ZM22 15v7h-7v4h7v7h4v-7h7v-4h-7v-7h-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHealthNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}