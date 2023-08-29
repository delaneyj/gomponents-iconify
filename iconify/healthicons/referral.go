package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Referral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 6a2 2 0 0 0-2 2v1H6v19h22V9h-5V8a2 2 0 0 0-2-2h-8Zm8 12a2 2 0 0 0 2-2v-1h3v11h-5v-6h-8v6H8V15h3v1a2 2 0 0 0 2 2h8ZM8 11h3v2H8v-2Zm18 2h-3v-2h3v2Zm-7 9v4h-4v-4h4ZM16 8v3h-3v2h3v3h2v-3h3v-2h-3V8h-2Z" clip-rule="evenodd"/><path d="M30 27a4 4 0 0 1 4-4a4 4 0 0 1 4 4a4 4 0 0 1-4 4a4 4 0 0 1-4-4Zm4 6c-2.67 0-8 1.462-8 4.364V40h-9c-.729 0-1.202-.263-1.503-.602C15.18 39.041 15 38.537 15 38v-3.586l3.293 3.293l1.414-1.414L14 30.586l-5.707 5.707l1.414 1.414L13 34.414V38c0 .963.32 1.959 1.003 2.727C14.702 41.513 15.729 42 17 42h25v-4.636C42 34.462 36.67 33 34 33Z"/></g>`),
		g.Group(children),
	)
}