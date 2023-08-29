package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreatPyramidOfGiza(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f4aa41" d="m50.846 56l-9.8-10.557l.352-1.979l2.423-2.892l2.298-1.182l4.244-5.064L57.431 56h-6.585z"/><path fill="#e27022" d="M51.957 55L25.968 26.963L35.078 56h5.306l11.573-1z"/><path fill="#f4aa41" d="m33.477 35.064l-7.509-8.101l2.642 8.101h4.867z"/><path fill="#f4aa41" d="m2.473 55l23.495-28.037L35.111 55H2.473z"/><path fill="#fcea2b" d="m19.18 35.064l6.788-8.101l2.642 8.101h-9.43z"/><path fill="#e27022" d="M69.527 55L50.363 34.326L57.105 55h12.422z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m44.586 41.219l5.777-6.893L57.105 55"/><path d="M69.527 55L50.363 34.326L57.105 55h12.422zM2.473 55l23.495-28.037L35.111 55H2.473zm36.976 0h12.508L25.968 26.963"/></g>`),
		g.Group(children),
	)
}