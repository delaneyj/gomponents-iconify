package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.81 43.147l9.948-12.673a.923.923 0 0 0-.726-1.492H17.138a.923.923 0 0 0-.726 1.492l9.947 12.673a.923.923 0 0 0 1.451 0ZM26.651 18.15l-5.938 7.567a.55.55 0 0 0 .433.89h11.877a.55.55 0 0 0 .433-.89l-5.938-7.566a.55.55 0 0 0-.867 0Zm-9.684 9.19l5.938-7.566a.55.55 0 0 0-.433-.89H10.596a.55.55 0 0 0-.434.89l5.939 7.566a.55.55 0 0 0 .866 0Zm9.798-17.314l-4.385 5.587a.407.407 0 0 0 .32.658h8.77c.339 0 .53-.391.32-.658l-4.385-5.587a.407.407 0 0 0-.64 0Zm.32-5.526l-3.078 3.921h6.155L27.085 4.5zm-10.482 8.275l-3.477 4.429h6.954l-3.477-4.429zm0-4.541l-2.473 3.151h4.946l-2.473-3.151zm0-3.734l-1.84 2.345h3.68L16.603 4.5z"/>`),
		g.Group(children),
	)
}