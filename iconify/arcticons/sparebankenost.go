package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparebankenost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.807 8.421c-3.714.262-8.447 2.344-11.234 10.6c-3.693 10.952-3.89 19.094-11.89 24.766A21.474 21.474 0 0 0 38.807 8.42ZM24 2.5A21.434 21.434 0 0 0 6.88 36.917c6.984-5.224 6.699-5.406 12.339-17.125C24.016 9.822 30.979 6.864 36.65 6.62A21.489 21.489 0 0 0 24 2.5Zm0 0"/>`),
		g.Group(children),
	)
}