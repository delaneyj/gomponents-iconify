package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Zm6 3.25a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5H8Zm-.414-8.664C8.172 6 9.114 6 11 6h2c1.886 0 2.828 0 3.414.586C17 7.172 17 8.114 17 10v.25h2a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h2V10c0-1.886 0-2.828.586-3.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}