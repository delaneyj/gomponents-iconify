package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M55.785 34.99a1 1 0 0 1 .105 1.041l-3.894 8.063a1.002 1.002 0 0 1-.706.546l-10.798 2.138a1.007 1.007 0 0 1-.794-.18L16.214 29.02a1 1 0 0 1-.176-1.432l2.438-2.997a.998.998 0 0 1 .326-.262l18.305-9.222a1.015 1.015 0 0 1 .658-.085l3.819.807a.997.997 0 0 1 .637.442l2.673 4.144l1.62 2.49Z"/><path fill="#3f3f3f" d="M26.692 56h20.096l-6.471-10.181l-23.374-17.6l-1.129 4.827L26.692 56z"/><path fill="#9b9b9a" d="m54.78 36.031l1.348 11.513L54.469 56h-7.681l-6.471-10.181l10.622-1.879l3.841-7.909z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m45.797 23.608l9.193 11.989l1.138 11.642L54.469 56H26.692L16 32.86l.813-4.641l2.439-2.997L37.558 16l3.819.807l2.643 4.11l-5.012 3.353"/><path d="m47.359 56l-7.061-10.203l-23.485-17.578m38.177 7.378l-3.895 8.062l-10.797 2.138"/></g>`),
		g.Group(children),
	)
}