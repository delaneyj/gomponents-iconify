package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonBikingDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#6a462f" stroke="#6a462f"><circle cx="34.386" cy="11.063" r="2.969"/><path d="M34.386 20.925L36.417 18l3.833.084l6.207 14.15L44.417 36l-2.527 2.514l-3.12 11.154l-3.406 6.75l-2.404-2.037l.919-7.078l-3.771-6.977l.097-3.326l8.045-5l-2.886-7.5"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="34.386" cy="11.063" r="2.969"/><path stroke-linecap="round" stroke-linejoin="round" d="m26.333 28.083l4.27-2.947c.906-.625 2.16-1.875 2.788-2.779l1.885-2.714c.627-.904 2.003-1.624 3.058-1.601c1.054.023 2.277.866 2.72 1.873l4.6 10.487c.442 1.008.374 2.623-.15 3.591l-.134.248c-.524.968-1.78 2.114-2.79 2.55l-5.492 2.363a2.04 2.04 0 0 0-1.071 2.639l2.634 6.36c.421 1.016.091 2.297-.734 2.847c-.825.55-1.928.208-2.451-.76l-4.406-8.154c-.524-.968-.93-2.508-.904-3.423c.028-.915.814-2.138 1.747-2.719l4.649-2.888c.933-.581 1.48-1.929 1.213-2.996L36.25 24"/><circle cx="20.084" cy="52" r="10.166"/><circle cx="52.084" cy="52" r="10.166"/><path stroke-linecap="round" stroke-linejoin="round" d="m20.25 52l2-24H26"/><path d="m37.978 50.959l-.984 3.133c-.33 1.05-1.42 1.908-2.422 1.908c-1.002 0-1.707-.893-1.565-1.983l.869-6.718m8.531-10.434L38.77 48.44"/></g>`),
		g.Group(children),
	)
}