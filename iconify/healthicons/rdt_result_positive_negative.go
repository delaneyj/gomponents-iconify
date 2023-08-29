package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultPositiveNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsRdtResultPositiveNegative0)"><path d="M18 40a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-1.5-7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill-rule="evenodd" d="M14 12a4 4 0 0 1 8 0v13a4 4 0 0 1-8 0V12Zm6 0v2h-4v-2a2 2 0 1 1 4 0Zm-4 9v-5h4v5h-4Zm0 2v2a2 2 0 1 0 4 0v-2h-4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm12 4a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H12Zm18.85 4.105c.238-.105.495-.131.743-.079c.248.053.485.185.677.39a1 1 0 1 0 1.46-1.368a3.301 3.301 0 0 0-1.722-.978a3.207 3.207 0 0 0-1.966.206a3.383 3.383 0 0 0-1.495 1.304A3.62 3.62 0 0 0 28 9.5c0 .678.188 1.346.547 1.92c.36.574.877 1.031 1.495 1.304a3.208 3.208 0 0 0 1.966.206a3.3 3.3 0 0 0 1.722-.978a1 1 0 1 0-1.46-1.368c-.192.205-.43.337-.677.39a1.207 1.207 0 0 1-.742-.079a1.384 1.384 0 0 1-.609-.537A1.621 1.621 0 0 1 30 9.5c0-.31.087-.61.242-.858c.156-.248.37-.432.609-.537ZM28.5 15a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-1v4a1 1 0 1 1-2 0v-4h-1a1 1 0 0 1-1-1ZM30 27.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .472.333a1 1 0 1 0 1.886-.666A2.5 2.5 0 0 0 31.5 25h-1a2.5 2.5 0 0 0 0 5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.472-.333a1 1 0 1 0-1.886.666A2.5 2.5 0 0 0 30.5 33h1a2.5 2.5 0 0 0 0-5h-1a.5.5 0 0 1-.5-.5Zm.5 8.5a.5.5 0 0 0 0 1h1a2.5 2.5 0 0 1 0 5h-1a2.5 2.5 0 0 1-2.358-1.667a1 1 0 1 1 1.886-.666a.5.5 0 0 0 .472.333h1a.5.5 0 0 0 0-1h-1a2.5 2.5 0 0 1 0-5h1a2.5 2.5 0 0 1 2.358 1.667a1 1 0 1 1-1.886.666A.5.5 0 0 0 31.5 36h-1Zm7.5-9.5a1 1 0 0 0-1.721-.692l-1 1.041A1 1 0 0 0 36 28.542V31.5a1 1 0 1 0 2 0v-5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRdtResultPositiveNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}