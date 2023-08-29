package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PleadingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="m35.94 12.58l-8.9 1.73l-5.46 3.12l-5.04 5.23L13 30.24v12.72l5.77 9.69l8.27 5.62l8.9 1.73l7.4-1.19l7.94-4.46L56.55 48l3.09-9.49l-.77-8.23l-3.93-8.15l-4.34-4.46l-7.47-3.97z"/><circle cx="44.45" cy="31.38" r="6.5" fill="#fff"/><circle cx="27.76" cy="31.38" r="6.5" fill="#fff"/><circle cx="36" cy="36" r="23" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29.02 46.68c1.379-1.678 4.279-2.713 7.347-2.628c2.817.079 5.358 1.091 6.612 2.628"/><circle cx="27.56" cy="31.38" r="6.5" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><circle cx="44.65" cy="31.38" r="6.5" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"/><path fill-rule="evenodd" d="M27.53 25.77a4.619 4.619 0 0 0-4.592 4.619a4.619 4.619 0 0 0 4.619 4.619a4.619 4.619 0 0 0 4.619-4.619a4.619 4.619 0 0 0-4.619-4.619a4.619 4.619 0 0 0-.027 0zm-1.453 1.828a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2zM44.62 25.77a4.619 4.619 0 0 0-4.592 4.619a4.619 4.619 0 0 0 4.619 4.619a4.619 4.619 0 0 0 4.619-4.619a4.619 4.619 0 0 0-4.619-4.619a4.619 4.619 0 0 0-.027 0zm-1.453 1.828a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2z" paint-order="stroke fill markers"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.31 18.02c1.185 2.421 3 5 9 4m-19.35-4c-1.185 2.421-3 5-9 4"/>`),
		g.Group(children),
	)
}