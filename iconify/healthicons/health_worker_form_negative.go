package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthWorkerFormNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsHealthWorkerFormNegative0)"><path d="M20 6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-8Zm-4 14a1 1 0 1 0 0 2h7a1 1 0 1 0 0-2h-7Zm-1-4a1 1 0 0 1 1-1h15.5a1 1 0 1 1 0 2H16a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M15 28a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4Zm2 3v-2h2v2h-2Z" clip-rule="evenodd"/><path d="M28 32a3 3 0 1 0 0-6a3 3 0 1 0 0 6Zm-6 4.5c0-2.116 3.997-3.182 6-3.182s6 1.066 6 3.182V39H22v-2.5Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm17 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h4a3 3 0 0 1 3 3v31a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V10a3 3 0 0 1 3-3h4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHealthWorkerFormNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}