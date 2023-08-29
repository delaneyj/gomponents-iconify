package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.906 3C15.29 3 14 4.23 14 5.656c0 1.336.469 2.328.938 2.969c.292.402.421.469.624.625l.188.75H5c-1.645 0-3 1.355-3 3s1.355 3 3 3h3.563l1.656 7.625A3.016 3.016 0 0 0 13.156 26H30V10h-5.594l-6.781-6.719L17.312 3zm-.312 2.094L23 11.406V24h-9.844c-.476 0-.898-.313-1-.781l-1.781-8.438l-.188-.781H5c-.566 0-1-.434-1-1c0-.566.434-1 1-1h13.344l-.313-1.25l-.593-2.25l-.125-.406l-.344-.188s-.18-.086-.438-.437c-.258-.352-.531-.91-.531-1.813c0-.308.172-.48.594-.562zM25 12h3v12h-3z"/>`),
		g.Group(children),
	)
}