package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwabPliers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<ellipse cx="31.584" cy="51.77" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" rx="3.729" ry="3.003" transform="rotate(-56.229 31.584 51.77)"/><path d="m30.583 48.93l-.3-2.169s-1.138.02-.587-.784l2.361-.28l.34 2.46a3.2 3.2 0 0 0-1.814.773Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m31.221 46.153l6.867-13.439"/><ellipse cx="23.568" cy="47.537" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" rx="3.729" ry="3.003" transform="rotate(-62.651 23.568 47.537)"/><path d="m26.202 46.375l2.043-.785s.532 1.006.971.135l-.8-2.098l-2.316.89a3.172 3.172 0 0 1 .102 1.858Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m28.34 44.648l8.953-12.761m9.73-12.61s4.32 2.918-.526 5.98l-5.893 3.108M44.32 18.11s-4.762-2.123-4.797 3.61l.481 6.645"/><rect width="3.471" height="5.679" x="37.289" y="27.319" rx="1.736" transform="rotate(30 39.025 30.157)"/><circle cx="39.02" cy="30.132" r=".613" fill="#fff"/>`),
		g.Group(children),
	)
}