package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><g stroke="#000" stroke-linejoin="round" stroke-width="2"><path fill="none" d="M5 17h62v38H5z"/><path fill="none" stroke-linecap="round" d="M9 21v31m3-31v31m0-31v31m8-31v29m8-29v29"/><path stroke-linecap="round" d="M15 50V21h2v29h-2Zm8 0V21h2v29h-2Zm8 0V21h1v29h-1Z"/><path fill="none" stroke-linecap="round" d="M46 21v29m3-29v29m8-29v29"/><path stroke-linecap="round" d="M41 50V21h2v29h-2Zm11 0V21h2v29h-2Z"/><path fill="none" stroke-linecap="round" d="M60 21v31m3-31v31M35 21v31m3-31v31"/></g>`),
		g.Group(children),
	)
}