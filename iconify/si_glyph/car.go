package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<defs><path id="siGlyphCar0" d="M6.514 1.607C6.319 2.141 5.934 3 5.934 3h3.092V.957H7.544c-.28 0-.843.131-1.03.65zm5.932-.64H9.968V3h4.086l-.484-1.447s-.41-.586-1.124-.586z"/></defs><g fill="none" fill-rule="evenodd" transform="translate(1 4)"><circle cx="12.49" cy="6.49" r="1.49" fill="currentColor"/><circle cx="3.49" cy="6.49" r="1.49" fill="currentColor"/><use href="#siGlyphCar0"/><use href="#siGlyphCar0"/><path fill="currentColor" d="M16 4.836c0-.597-.078-1.137-.753-1.524l-.585-1.974C14.662.947 13.747 0 12.618 0H7.339C6.21 0 5.296 1.088 5.296 1.338L4.432 3l-.781.107C1.633 3.107.583 4.334.583 5.515L.024 6.984h1.007a2.454 2.454 0 0 1-.056-.518a2.508 2.508 0 0 1 2.519-2.498a2.508 2.508 0 0 1 2.518 2.498c0 .178-.02.351-.056.518h4.074a2.454 2.454 0 0 1-.056-.518a2.508 2.508 0 0 1 2.519-2.498a2.508 2.508 0 0 1 2.519 2.498c0 .178-.02.351-.056.518H16V4.836zM9.025 3H5.933s.385-.858.58-1.393c.188-.52.751-.65 1.03-.65h1.482V3zm.944 0V.967h2.478c.714 0 1.124.586 1.124.586L14.055 3H9.969z"/></g>`),
		g.Group(children),
	)
}