package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandOpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHandOpenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M6.938 16.94a1.085 1.085 0 0 0 .037.07c.18.314.354.63.525.94l.003.005c.234.425.462.839.68 1.204c.393.658.83 1.292 1.4 1.827c1.21 1.134 2.801 1.625 5.242 1.625c2.28 0 3.884-1.149 4.863-2.745c.946-1.542 1.312-3.486 1.312-5.26V9.003a2.401 2.401 0 0 0-2.516-2.399V6.26a2.401 2.401 0 0 0-3.296-2.23a2.398 2.398 0 0 0-3.56.355A2.401 2.401 0 0 0 8.65 6.717v3.793a2.37 2.37 0 0 0-3.4 3.057l1.687 3.373l.001.002Zm4.685-10.224a.572.572 0 1 0-1.143 0v7.875a.684.684 0 0 1-.51.678a.682.682 0 0 1-.525-.072a.69.69 0 0 1-.271-.297l-1.318-2.636a.541.541 0 0 0-.969.484l1.682 3.364c.207.361.387.69.555.995c.22.4.42.763.63 1.114c.359.6.692 1.066 1.08 1.43c.74.693 1.797 1.13 3.991 1.13c1.56 0 2.61-.742 3.303-1.872c.713-1.162 1.042-2.74 1.042-4.303V9.003a.572.572 0 0 0-1.143 0v3.989a.686.686 0 1 1-1.372 0V6.259a.572.572 0 0 0-1.144 0v6.06a.686.686 0 0 1-1.372 0V5.801a.572.572 0 0 0-1.143 0v6.976a.686.686 0 0 1-1.373 0v-6.06Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHandOpenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}