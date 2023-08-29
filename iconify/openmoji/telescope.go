package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="m45.67 22.205l3.746-1.33c-.071.016-.143.038-.214.038c-.41 0 2.972-2.616 2.972-2.616l-3.37-7.737s-4.186 1.287-3.945.99l-3.143 1.63c.238-.108.479 1.058.59 1.311l2.475 6.423c.109.253 1.109 1.188.888 1.291z"/><path fill="#9b9b9a" d="m29.095 28.198l3.828-1.415a5860.38 5860.38 0 0 1 12.405-5.809l-3.2-6.893s-11.85 4.906-11.66 4.809l1.681-.855l-5.226 2.606l2.172 7.557z"/><path fill="#9b9b9a" d="m14.659 27.983l1.76 3.933s9.08-2.49 8.95-2.443l-1.144.414l5.403-1.89l-2.44-7.488l-4.024 2.138c.237-.1-8.506 5.336-8.506 5.336z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m16.019 56.662l13.13-20.623h2.475l12.271 20.617M30.462 36.741v23.516M9.748 32.575l2.087-.863m33.798-19.533l3.17-1.619l3.578 8.212l-3.179 1.141m-22.38.011l3.531 11.437"/><circle cx="30.353" cy="31.361" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m25.028 28.533l-8.831 3.199l-1.538-3.558l8.887-4.617m7.376-3.776l11.206-5.699l3.127 7.228l-12.599 4.519"/>`),
		g.Group(children),
	)
}