package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icicles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 5v2.375l.813.281C7.546 8.27 8 9.156 8 10.125V21c0 1.645 1.355 3 3 3s3-1.355 3-3v-2.125c0-.73.66-1.195 1.344-.938a1 1 0 0 1 .656.938V25c0 1.645 1.355 3 3 3s3-1.355 3-3v-8.969c0-1.273.98-2.234 2.188-2.281h.03l1.782.625V12zm3.813 3.719L21.5 12.813c-.906.78-1.5 1.917-1.5 3.218V25c0 .566-.434 1-1 1c-.566 0-1-.434-1-1v-6.125c0-1.246-.77-2.375-1.938-2.813c-1.93-.722-4.062.75-4.062 2.813V21c0 .566-.434 1-1 1c-.566 0-1-.434-1-1V10.125c0-.48-.063-.953-.188-1.406z"/>`),
		g.Group(children),
	)
}