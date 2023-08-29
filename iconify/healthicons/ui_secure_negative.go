package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiSecureNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsUiSecureNegative0)"><path d="M18 18h12v-2a6 6 0 0 0-12 0v2Zm-3 8h18v-2H15v2Zm18 5H15v-2h18v2Zm-18 5h18v-2H15v2Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm11 18a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h26a2 2 0 0 0 2-2V20a2 2 0 0 0-2-2h-3v-2c0-5.523-4.477-10-10-10s-10 4.477-10 10v2h-3Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUiSecureNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}