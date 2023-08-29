package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Om(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12h48v47.83H12z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="3" d="M25.24 30.12c12.84-6.27 12.38 4.635 4.133 7.694c0 0 6.959.88 6.9 5.387c-.059 4.506-2.585 7.392-6.788 7.13c-2.878-.263-7.598-1.336-9.064-8.283"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="3" d="M35.37 24.17s5.631 7.013 11.78.896M34 39.48s5.35 1.423 6.959-2.687c0 0 2.813-6.988 7.322-4.92s3.465 10.09 2.277 11.98c0 0-4.101 7.366-9.378 1.375"/><circle cx="40.11" cy="22.15" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h48v48H12z"/>`),
		g.Group(children),
	)
}