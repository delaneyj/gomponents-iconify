package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CholeraNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsCholeraNegative0)"><path d="M25.026 33.23a1 1 0 1 0 2-.024A1.192 1.192 0 0 1 28.204 32a1 1 0 1 0-.024-2a3.192 3.192 0 0 0-3.154 3.23ZM31 32a1 1 0 0 1 1-1a3 3 0 0 1 3 3a1 1 0 1 1-2 0a1 1 0 0 0-1-1a1 1 0 0 1-1-1Zm-2.174 8.118a1 1 0 1 0-.43-1.953a1.192 1.192 0 0 1-1.42-.91a1 1 0 1 0-1.953.43a3.192 3.192 0 0 0 3.803 2.433ZM35 38a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-4-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm24 4H10v8h15a1 1 0 0 1 1 1v1h8c0-5.523-4.477-10-10-10Zm12 12H24v4h12v-4Zm-6 28a8 8 0 0 0 8-8c0-7-8-14-8-14s-8 7-8 14a8 8 0 0 0 8 8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCholeraNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}